import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface Ssd extends BaseProduct {
    type: string; // Тип SSD (например, SATA, NVMe)
    capacity: number; // Объем SSD (например, 500GB, 1TB)
    interface: string; // Интерфейс подключения SSD (например, SATA III, PCIe Gen 3)
    memoryType: string; // Тип памяти (например, TLC, QLC)
    read: number; // Скорость чтения (например, 550 MB/s)
    write: number; // Скорость записи (например, 520 MB/s)
    formFactor: string; // Форм-фактор SSD (например, 2.5", M.2)
    mftb: number; // MTBF (среднее время между отказами) в часах
    size: number[]; // Размеры SSD в миллиметрах [длина, ширина, высота]
    weight: number; // Вес SSD в граммах
}
