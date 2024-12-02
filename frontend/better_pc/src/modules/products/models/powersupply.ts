import {BaseProduct} from "@/modules/products/models/product_common"; // Общий интерфейс General

// Описание структуры подключения (SATA, Molex, PCI-E)
export interface Connectors {
    sata?: number; // Количество портов SATA
    molex?: number; // Количество разъемов Molex
    pciE?: number[]; // Список слотов PCI-E (например, [6, 8] для двух разъемов PCI-E 6-pin и 8-pin)
}

// Описание структуры питания процессора
export interface CpuPower {
    amount?: number; // Количество разъемов питания процессора
    type?: number[]; // Типы разъемов (например, [4, 8] для разъемов 4-pin и 8-pin)
}

export interface PowerSupply extends BaseProduct {
    formFactor?: string; // Форм-фактор блока питания (например, ATX, SFX)
    outputPower?: number; // Выходная мощность блока питания (в ваттах)
    connectors: Connectors; // Разъемы блока питания (SATA, Molex, PCI-E)
    modules?: boolean; // Наличие модульных кабелей
    mbPower?: number; // Питание материнской платы (в пинах)
    cpuPower?: CpuPower; // Питание процессора, включая количество и тип разъемов
}
