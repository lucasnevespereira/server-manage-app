type SortTypeSelectProps = {
    value: SortType;
    onChange: (event: React.ChangeEvent<HTMLSelectElement>) => void;
};

export default function SortTypeSelect({value, onChange}: SortTypeSelectProps) {
    return (
        <select
            className="select select-bordered"
            value={value}
            onChange={onChange}
            data-testid="type-select"
        >
            <option value="name">Name</option>
            <option value="type">Type</option>
            <option value="status">Status</option>
        </select>
    );
}