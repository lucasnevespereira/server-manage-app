import {useFetchApi} from "./hooks/useFetchApi.ts";
import "./App.css"
import {useState} from "react";
import ServerList from "./components/ServerList";
import SortOrderSelect from "./components/SortOrderSelect";
import SortTypeSelect from "./components/SortTypeSelect";

function App() {
    const {data, isFetching} = useFetchApi<ServerData>('/servers/');
    const [sortOrder, setSortOrder] = useState<SortOrder>("asc");
    const [sortType, setSortType] = useState<SortType>("name");

    const onSortTypeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selected = event.target.value as SortType;
        setSortType(selected);
    };

    const onSortOrderChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selected = event.target.value as SortOrder;
        setSortOrder(selected);
    };

    const sortedServers = data?.servers?.slice().sort((a: Server, b: Server) => {
        const compareResult = a[sortType].localeCompare(b[sortType])
        return sortOrder === "asc" ? compareResult : -compareResult;
    });

    return (
        <div className="container mx-auto p-6">
            <div className="flex justify-between">
                <h2 className="text-2xl text-primary font-bold mb-6">Server Manage</h2>
                <div className="flex gap-2">
                    <SortTypeSelect value={sortType} onChange={onSortTypeChange}/>
                    <SortOrderSelect value={sortOrder} onChange={onSortOrderChange}/>
                </div>
            </div>
            <div className="grid grid-cols-1 gap-4">
                {isFetching ?
                    <div className="flex p-5">
                        <span className="animate-bounce text-primary text-2xl">Loading...</span>
                    </div>
                    :
                    <ServerList servers={sortedServers || []}/>
                }
            </div>
        </div>
    )
}

export default App
