type ServerData = {
    servers: Server[]
    total: number
}

type Server = {
    id: string
    name: string
    type: string
    status: string
}

type SortOrder = "asc" | "desc";
type SortType = "name" | "type" | "status"