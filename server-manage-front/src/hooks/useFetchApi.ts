import {useEffect, useState} from "react";
import axios from "axios";

export function useFetchApi<T = unknown>(endpoint: string) {
    const [data, setData] = useState<T | null>(null)
    const [isFetching, setIsFetching] = useState(true)
    const baseUrl: string = import.meta.env.VITE_API_BASE_URL;
    const url = baseUrl + endpoint

    useEffect(() => {
        void axios.get(url)
            .then(response => {
                setData(response.data)
            })
            .finally(() => {
                setIsFetching(false)
            })
    }, [])

    return {data, isFetching}
}