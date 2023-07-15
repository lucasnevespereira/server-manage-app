import {render, screen} from '@testing-library/react';
import App from './App';
import ServerList from "./components/ServerList";
import SortOrderSelect from "./components/SortOrderSelect";
import SortTypeSelect from "./components/SortTypeSelect";

type MockResponseData = {
    data: ServerData
    isFetching: boolean
}
const mockResponseData: MockResponseData = {
    data: {
        servers: [
            {id: '1', name: 'Server 1', type: 'small', status: 'running'},
            {id: '2', name: 'Server 2', type: 'large', status: 'stopped'},
        ],
        total: 2,
    },
    isFetching: false
}
const mockOnChange = () => void {}

describe('App', () => {
    test('renders the app title', () => {
        render(<App/>);
        const titleElement = screen.getByText('Server Manage');
        expect(titleElement).toBeInTheDocument();
    });

    test('renders "No servers found" message when no servers are provided', () => {
        render(<ServerList servers={[]}/>);
        const noServersMessage = screen.getByText('No servers found');
        expect(noServersMessage).toBeInTheDocument();
    });

    test('renders the list of servers when servers are provided', () => {
        const mockServers = mockResponseData.data.servers
        render(<ServerList servers={mockServers}/>);
        const serverElements = screen.getAllByTestId('server-card');
        expect(serverElements).toHaveLength(mockServers.length);
    });

    test('renders the order select element with options', () => {
        render(<SortOrderSelect value="asc" onChange={mockOnChange}/>);

        const selectElement = screen.getByTestId('order-select');
        expect(selectElement).toBeInTheDocument();

        const options = screen.getAllByRole('option');
        expect(options).toHaveLength(2);
        expect(options[0]).toHaveValue('asc');
        expect(options[1]).toHaveValue('desc');
    });

    test('renders the type select element with options', () => {
        render(<SortTypeSelect value="name" onChange={mockOnChange}/>);

        const selectElement = screen.getByTestId('type-select');
        expect(selectElement).toBeInTheDocument();

        const options = screen.getAllByRole('option');
        expect(options).toHaveLength(3);
        expect(options[0]).toHaveValue('name');
        expect(options[1]).toHaveValue('type');
        expect(options[2]).toHaveValue('status');
    });
});