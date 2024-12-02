import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface DriveBays {
    d35: number; // Количество отсеков для 3.5" накопителей
    d25: number; // Количество отсеков для 2.5" накопителей
}

export interface Housing extends BaseProduct {
    formFactor: string; // Форм-фактор корпуса (например, "ATX", "mATX", "ITX")
    driveBays: DriveBays; // Объект с информацией о количестве отсеков для накопителей
    mbFormFactor: string; // Поддерживаемые форм-факторы материнских плат
    psFormFactor: string; // Форм-фактор блока питания (например, "ATX", "SFX")
    expansionSlots: number; // Количество слотов расширения
    graphicCardSize: number; // Максимальная длина видеокарты (в мм)
    coolerHeight: number; // Максимальная высота процессорного кулера (в мм)
    size: number[]; // Размеры корпуса (длина, ширина, высота, в мм)
    weight: number; // Вес корпуса (в кг)
}
