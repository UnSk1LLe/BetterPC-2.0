import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

export interface RamMb {
    slots: number; // Количество слотов для оперативной памяти
    type: string; // Тип оперативной памяти (например, DDR4, DDR5)
    maxFrequency: number; // Максимальная частота памяти (в МГц)
    maxCapacity: number; // Максимальный поддерживаемый объем памяти (в ГБ)
}

export interface Interfaces {
    sata3: number; // Количество портов SATA3
    m2: number; // Количество слотов M.2
    pciE1x: number; // Количество слотов PCI-E 1x
    pciE16x: number; // Количество слотов PCI-E 16x
}

export interface Motherboard extends BaseProduct {
    socket?: string; // Тип сокета процессора (например, AM4, LGA1700)
    chipset?: string; // Название чипсета (например, B550, Z790)
    formFactor?: string; // Форм-фактор материнской платы (например, ATX, mATX, Mini-ITX)
    ram: RamMb; // Характеристики оперативной памяти
    interfaces: Interfaces; // Интерфейсы подключения
    pciStandard?: number; // Версия PCI Express (например, 3, 4, 5)
    mbPower?: number; // Питание материнской платы (в пинах)
    cpuPower?: number; // Питание процессора (в пинах)
}

