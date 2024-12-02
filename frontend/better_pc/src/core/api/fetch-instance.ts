export type RequestOptions = {
    headers?: Record<string, string>
    params?: Record<string, string | number> | URLSearchParams
    body?: unknown
}

class FetchWrapper {
    baseUrl: string =
        process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080'

    buildUrlWithParams(
        url: string,
        params?: Record<string, string | number> | URLSearchParams
    ): string {
        if (!params) return url

        let queryString: string

        if (params instanceof URLSearchParams) {
            queryString = params.toString()
        } else {
            queryString = new URLSearchParams(
                Object.entries(params).map(([key, value]) => [key, String(value)])
            ).toString()
        }

        return `${url}?${queryString}`
    }

    private async request<T = unknown>(
        method: 'GET' | 'POST' | 'PUT' | 'DELETE',
        url: string,
        options: RequestOptions = {}
    ): Promise<T> {
        const { headers, params, body } = options
        const fullUrl = this.buildUrlWithParams(`${this.baseUrl}${url}`, params)

        const isFormData = body instanceof FormData

        const response = await fetch(fullUrl, {
            method,
            credentials: 'include',
            cache: 'no-store',
            referrerPolicy: 'no-referrer',
            redirect: 'follow',
            headers: {
                ...(isFormData ? {} : { 'Content-Type': 'application/json' }),
                ...(headers || {})
            },
            body: body ? (isFormData ? body : JSON.stringify(body)) : undefined
        })

        if (!response.ok) {
            throw new Error(`Error ${response.status}: ${response.statusText}`)
        }

        const contentType = response.headers.get('Content-Type')
        if (!contentType || response.status === 204) {
            return { message: 'No body' } as T
        }

        try {
            if (contentType.includes('application/json')) {
                const result: T = await response.json()
                return result
            } else {
                throw new Error('Unsupported content type')
            }
        } catch (error) {
            throw new Error(`Failed to parse response: ${error}`)
        }
    }

    async get<T = unknown>(url: string, options?: RequestOptions): Promise<T> {
        return this.request<T>('GET', url, options)
    }

    async post<T = unknown>(
        url: string,
        data?: unknown,
        options?: RequestOptions
    ): Promise<T> {
        return this.request<T>('POST', url, { ...options, body: data })
    }

    async put<T = unknown>(
        url: string,
        data?: unknown,
        options?: RequestOptions
    ): Promise<T> {
        return this.request<T>('PUT', url, { ...options, body: data })
    }

    async delete<T = any>(url: string, options?: RequestOptions): Promise<T> {
        return this.request<T>('DELETE', url, options)
    }
}

export const fetchWrapper = new FetchWrapper()