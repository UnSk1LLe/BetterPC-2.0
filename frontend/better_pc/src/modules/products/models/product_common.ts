import {General} from "@/modules/products/models/general/general";
import {Cpu} from "@/modules/products/models/cpu";
import {Motherboard} from "@/modules/products/models/motherboard";
import {Ram} from "@/modules/products/models/ram";
import {Gpu} from "@/modules/products/models/gpu";
import {Ssd} from "@/modules/products/models/ssd";
import {Hdd} from "@/modules/products/models/hdd";
import {Cooling} from "@/modules/products/models/cooling";
import {PowerSupply} from "@/modules/products/models/powersupply";
import {Housing} from "@/modules/products/models/housing";

export interface BaseProduct {
    id?: string; // Уникальный идентификатор продукта
    general: General; // Общая информация о продукте
}

export type ProductType =
    | 'cpu'
    | 'motherboard'
    | 'ram'
    | 'gpu'
    | 'ssd'
    | 'hdd'
    | 'cooling'
    | 'powersupply'
    | 'housing';

