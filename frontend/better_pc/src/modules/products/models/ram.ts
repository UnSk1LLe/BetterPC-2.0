import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface Ram extends BaseProduct {
    capacity: number; // Объем памяти (в гигабайтах)
    number: number; // Количество модулей
    formFactor: string; // Форм-фактор RAM (например, DIMM, SO-DIMM)
    rank: number; // Ранг памяти (например, одиночный или двухканальный)
    type: string; // Тип памяти (например, DDR4, DDR5)
    frequency: number; // Частота работы памяти (например, 3200 MHz)
    bandwidth: number; // Пропускная способность памяти (например, 25.6 GB/s)
    casLatency: string; // CAS латентность (например, 16-18-18)
    timingScheme: number[]; // Тайминги памяти в виде массива (например, [16, 18, 18, 36])
    voltage: number; // Напряжение работы памяти (например, 1.35V)
    cooling: string; // Охлаждение памяти (например, с радиатором, без охлаждения)
    height: number; // Высота модуля памяти (например, в миллиметрах)
}
