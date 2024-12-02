export interface LoginDTO {
    email: string
    password: string
}

export interface RegisterDTO {
    name: string
    surname: string
    email: string
    dob: string
    password: string
}

export enum Tokens {
    ACCESS = 'accessToken',
    REFRESH = 'refreshToken'
}