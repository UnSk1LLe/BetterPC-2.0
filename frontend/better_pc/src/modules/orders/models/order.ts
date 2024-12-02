import {ProductType} from "@/modules/products/models/product_common";

export interface Order {
    id: string; // Unique identifier for the order
    productList: Record<ProductType, ProductHeader[]>; // Map-like structure for products
    userId: string; // User identifier
    price: number; // Total price of the order
    status: OrderStatus; // Enum or string representation of the order status
    payment: PaymentDetails; // Payment information
    refunds?: RefundDetails[]; // Optional array of refunds
    createdAt: string; // ISO string for creation timestamp
    updatedAt: string; // ISO string for update timestamp
}

export interface PaymentDetails {
    paymentIntentId?: string; // Optional, similar to `omitempty` in Go
    isPaid: boolean;
}

export interface RefundDetails {
    refundId: string;
    amount: number; // In cents or lowest currency unit
    currency: string;
    refundedAt: string; // ISO string format for dates
}

export interface ProductHeader {
    id: string;
    price: number; // Price in lowest currency unit (e.g., cents)
    selectedAmount: number;
}

// Example OrderStatus enum
export enum OrderStatus {
    Created = 'CREATED',
    Pending = 'PENDING',
    Delivered = 'DELIVERED',
    Closed = 'CLOSED',
    Cancelled = 'CANCELLED',
}