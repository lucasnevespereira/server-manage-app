import ServerCard from "./ServerCard";

type ServerListProps = {
    servers: Server[]
};

export default function ServerList({servers}: ServerListProps) {
    return (
        <div className="grid grid-cols-1 gap-4" data-testid="server-list">
            {servers.length === 0 ? (
                <p className="text-xl font-bold">No servers found</p>
            ) : (
                servers.map((server) => (
                    <ServerCard server={server} key={server.id}/>
                ))
            )}
        </div>
    );
}