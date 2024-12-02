import {General} from "@/modules/products/models/general/general";
import {BaseProduct} from "@/modules/products/models/product_common";

export interface Cpu extends BaseProduct {
    main: MainCpu;
    cores?: CoresCpu;
    clockFrequency?: ClockFrequencyCpu;
    ram?: RamCpu;
    tdp?: number;
    graphics?: string;
    pciE?: number;
    maxTemperature?: number;
}

export interface MainCpu {
    category?: string;
    generation?: string;
    socket?: string;
    year?: number;
}

export interface CoresCpu {
    pCores?: number;
    eCores?: number;
    threads?: number;
    technicalProcess?: number;
}

export interface ClockFrequencyCpu {
    pCores?: number[];
    eCores?: number[];
    freeMultiplier?: boolean;
}

export interface RamCpu {
    channels?: number;
    types?: RamCpuType[];
    maxCapacity?: number;
}

export interface RamCpuType {
    type?: string;
    maxFrequency?: number;
}