import {BaseProduct} from "@/modules/products/models/product_common";

export interface StandardizedProduct extends BaseProduct {
    name: string
    description: string
}