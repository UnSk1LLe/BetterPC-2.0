import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface Gpu extends BaseProduct {
    architecture: string; // Архитектура GPU (например, "Ampere", "RDNA 3")
    memory: MemoryGpu; // Память GPU
    gpuFrequency: number; // Частота графического процессора (в МГц)
    processSize: number; // Техпроцесс (в нанометрах)
    maxResolution: string; // Максимальное разрешение (например, "7680x4320")
    interfaces: InterfacesGpu[]; // Интерфейсы подключения (например, HDMI, DisplayPort)
    maxMonitors: number; // Максимальное количество мониторов
    cooling: CoolingGpu; // Система охлаждения GPU
    tdp: number; // TDP GPU
    tdpR: number; // Рекомендуемая мощность для TDP
    powerSupply: number[]; // Мощность разъемов питания (например, [6, 8] для 6+8 pin)
    slots: number; // Количество слотов (например, 2.5)
    size: number[]; // Габариты GPU (например, [300, 140, 50])
}

// Подмодели для Gpu
export interface MemoryGpu {
    capacity: number; // Объем памяти (в ГБ)
    type: string; // Тип памяти (например, GDDR6, HBM2)
    interfaceWidth: number; // Ширина шины (в битах)
    frequency: number; // Частота памяти (в МГц)
}

export interface InterfacesGpu {
    type: string; // Тип интерфейса (например, HDMI, DisplayPort)
    number: number; // Количество разъемов
}

export interface CoolingGpu {
    type: string; // Тип охлаждения (например, активное, пассивное)
    fanNumber: number; // Количество вентиляторов
}