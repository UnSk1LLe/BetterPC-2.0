import {BaseProduct} from "@/modules/products/models/product_common";

export interface Cooling extends BaseProduct {
    type: string; // Тип охлаждения (например, воздушное, жидкостное)
    sockets: string[]; // Совместимые сокеты
    fans: number[]; // Количество вентиляторов
    rpm: number[]; // Скорости вращения вентиляторов
    tdp: number; // Поддерживаемый уровень тепловыделения
    noiseLevel: number; // Уровень шума (в децибелах)
    mountType: string; // Тип крепления
    power: number; // Потребляемая мощность
    height: number; // Высота в мм
}