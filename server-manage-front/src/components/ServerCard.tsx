type ServerCardProps = {
    server: Server;
};

export default function ServerCard({server}: ServerCardProps) {
    const calculateBadge = (status: string) => {
        switch (status) {
            case "starting":
                return "badge badge-primary";
            case "running":
                return "badge badge-success";
            case "stopping":
                return "badge badge-warning";
            case "stopped":
                return "badge badge-error";
            default:
                return "badge badge-neutral";
        }
    };

    return (
        <div
            className="card border hover:cursor-pointer bg-base-100 shadow-lg rounded-lg p-4"
            key={server.id}
            data-testid="server-card"
        >
            <div className="flex justify-between items-center mx-3">
                <h2 className="text-xl font-bold mb-2">{server.name}</h2>
                <div className={calculateBadge(server.status)}>{server.status}</div>
            </div>

            <div tabIndex={0} className="collapse collapse-plus">
                <div className="collapse-title text-md font-medium">See Details</div>
                <div className="collapse-content">
                    <ul>
                        <li>
                            <b>ID:</b> {server.id}
                        </li>
                        <li>
                            <b>Name:</b> {server.name}
                        </li>
                        <li>
                            <b>Type:</b> {server.type}
                        </li>
                        <li>
                            <b>Status:</b> {server.status}
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    );
}