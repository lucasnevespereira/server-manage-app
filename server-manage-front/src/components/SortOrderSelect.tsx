type SortOrderSelectProps = {
    value: SortOrder;
    onChange: (event: React.ChangeEvent<HTMLSelectElement>) => void;
};

export default function SortOrderSelect({value, onChange}: SortOrderSelectProps) {
    return (
        <select
            className="select select-bordered"
            value={value}
            onChange={onChange}
            data-testid="order-select"
        >
            <option value="asc">Asc</option>
            <option value="desc">Desc</option>
        </select>
    );
}
