package cmd

import (
	"BetterPC_2.0/configs"
	jsonDecoders "BetterPC_2.0/datasets/decoders"
	"BetterPC_2.0/internal/repository/database/mongoDb"
	"BetterPC_2.0/internal/service"
	"BetterPC_2.0/pkg/data/models/products"
	"BetterPC_2.0/pkg/logging"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  `Database seeder`,
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	logger := logging.GetLogger()
	//initializing logger

	gin.ForceConsoleColor()

	logger.Infof("Start seed")

	err := configs.InitConfig() //initializing config path
	if err != nil {
		logger.Fatalf("failed to initialize config: %s", err.Error())
	}

	configs.SetConfig()

	mongoDb.MustConnectMongo(configs.GetConfig(), logger) //establishing connection to mongoDB database

	mongoDbConnection, err := mongoDb.GetMongoDB() //getting the established connection to mongoDb client and collections
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	productFilePaths := map[products.ProductType]string{
		products.ProductTypes.Cpu:         "./datasets/cpu.json",
		products.ProductTypes.Motherboard: "./datasets/motherboard.json",
		products.ProductTypes.Ram:         "./datasets/ram.json",
		products.ProductTypes.Gpu:         "./datasets/gpu.json",
		products.ProductTypes.Ssd:         "./datasets/ssd.json",
		products.ProductTypes.Hdd:         "./datasets/hdd.json",
		products.ProductTypes.Cooling:     "./datasets/cooling.json",
		products.ProductTypes.PowerSupply: "./datasets/powersupply.json",
		products.ProductTypes.Housing:     "./datasets/housing.json",
	}

	for productType, filePath := range productFilePaths {
		switch productType {
		case products.ProductTypes.Cpu:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeCpuList,
				logger,
			)
		case products.ProductTypes.Motherboard:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeMotherboardList,
				logger,
			)
		case products.ProductTypes.Ram:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeRamList,
				logger,
			)
		case products.ProductTypes.Gpu:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeGpuList,
				logger,
			)
		case products.ProductTypes.Ssd:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeSsdList,
				logger,
			)
		case products.ProductTypes.Hdd:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeHddList,
				logger,
			)
		case products.ProductTypes.Cooling:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeCoolingList,
				logger,
			)
		case products.ProductTypes.PowerSupply:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodePowerSupplyList,
				logger,
			)
		case products.ProductTypes.Housing:
			processAndInsertProducts(
				filePath,
				mongoDbConnection,
				productType,
				jsonDecoders.DecodeHousingList,
				logger,
			)
		// Add similar cases for other product types with their respective decoding functions.
		default:
			logger.Errorf("No decoding function implemented for %s", productType)
		}
	}

	logger.Info("Seed finished!")
}

func processAndInsertProducts[T products.Product](
	filePath string,
	mongoDbConnection mongoDb.Database,
	productType products.ProductType,
	decodeFunc func([]byte) ([]T, error),
	logger *logging.Logger,
) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open JSON file: %w", err))
	}
	defer file.Close()

	fileService := service.NewFileService(service.StaticFilesPath)

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed to read JSON file: %w", err))
	}

	// Decode the JSON data
	productList, err := decodeFunc(content)
	if err != nil {
		panic(fmt.Errorf("failed to decode JSON: %w", err))
	}

	// Process each product
	for _, product := range productList {
		imageURL := product.GetImage()
		imageFile, err := fetchImageAsFileHeader(imageURL)
		if err != nil {
			logger.Fatalf("failed to fetch image file: %s", err.Error())
		}

		imageName, err := fileService.AddProductImage(imageFile)
		if err != nil {
			logger.Errorf("Failed to process image for %s: %v\n", product.GetModel(), err)
			return
		}

		// Update image path
		product.SetImage(imageName)
	}

	// Convert to []*products.Product for insertion
	var productInterfaces []products.Product
	for _, product := range productList {
		productInterfaces = append(productInterfaces, product)
	}

	// Insert into the database
	mustInsertProducts(mongoDbConnection, productInterfaces, productType)
	logger.Infof("images from %s list have been fetched and inserted successfully.", productType)
}

func fetchImageAsFileHeader(imageURL string) (*multipart.FileHeader, error) {
	// Get the image from the URL
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: HTTP status %d", resp.StatusCode)
	}

	// Create a temporary file to hold the image data
	tempFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tempFile.Close()

	// Copy the image data to the temporary file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write image data: %v", err)
	}

	// Reopen the file for reading
	file, err := os.Open(tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to reopen temp file: %v", err)
	}
	defer file.Close()

	// Create a buffer and multipart.Writer to construct the FileHeader
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Create a multipart file field
	part, err := writer.CreateFormFile("file", filepath.Base(tempFile.Name()))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}

	// Copy the file content into the multipart field
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file content: %v", err)
	}
	writer.Close()

	// Parse the multipart content to get the FileHeader
	reader := multipart.NewReader(&buffer, writer.Boundary())
	form, err := reader.ReadForm(1024 * 1024) // Max memory usage 1 MB
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %v", err)
	}
	defer form.RemoveAll()

	// Return the first FileHeader
	for _, files := range form.File {
		if len(files) > 0 {
			return files[0], nil
		}
	}

	return nil, fmt.Errorf("no file found in multipart form")
}

func mustInsertProducts(db mongoDb.Database, products []products.Product, productType products.ProductType) {
	logger := logging.GetLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	docs := make([]interface{}, len(products))
	for i, p := range products {
		docs[i] = p
	}

	_, err := db.GetProductCollection(productType).InsertMany(ctx, docs)
	if err != nil {
		logger.Fatalf("failed to insert products: %s", err.Error())
	}
}
