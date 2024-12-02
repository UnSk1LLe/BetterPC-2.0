import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface Hdd extends BaseProduct {
    type: string; // Тип диска (например, "HDD", "SSHD")
    capacity: number; // Объем памяти (в ГБ)
    interface: string; // Интерфейс подключения (например, "SATA", "SAS")
    writeMethod: string; // Метод записи (например, "CMR", "SMR")
    transferRate: number; // Скорость передачи данных (в МБ/с)
    spindleSpeed: number; // Скорость вращения шпинделя (в RPM)
    formFactor: string; // Форм-фактор (например, "3.5", "2.5")
    mftb: number; // Среднее время наработки на отказ (в часах)
    size: number[]; // Размеры диска (длина, ширина, высота, в мм)
    weight: number; // Вес (в граммах)
}
