import {fetchWrapper} from "@/core/api/fetch-instance";
import {StandardizedProduct} from "@/modules/products/models/standardized";

class ProductsService {

    private readonly url: string = "/api/v1/shop/categories"

    async getStandardizedProducts(): Promise<StandardizedProduct[]> {
        try {
            return await fetchWrapper.get<StandardizedProduct[]>(`${this.url}/all`)
        } catch {
            return [] as StandardizedProduct[]
        }
    }
}