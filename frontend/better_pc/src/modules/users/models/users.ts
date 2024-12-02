export interface User {
    id: string
    name: string
    surname: string
    email: string
    dob: string
    isVerified: boolean
    role: Roles
    image: string
}

export enum Roles {
    USER = 'CUSTOMER',
    SHOP_ASSISTANT = 'SHOP_ASSISTANT',
    ADMIN = 'ADMIN'
}