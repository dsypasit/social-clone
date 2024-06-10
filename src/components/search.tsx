import React, {
  ChangeEvent,
  ChangeEventHandler,
  ReactEventHandler,
  useState,
} from "react";
import { Input } from "./ui/input";
import { Separator } from "./ui/separator";

interface Item {
  // Define the type of your list items here (optional)
  name: string; // Example property
}

const MockValue: Item[] = [
  { name: "aser" },
  { name: "ajajaj" },
  { name: "jkjkjkj" }, // Use objects if your items have properties
];

export function Search() {
  const [IsFocused, setIsFocused] = useState(false);
  const [filterValue, setFilterValue] = useState(""); // State for user input

  const filteredItems = React.useMemo(() => {
    if (!filterValue) {
      return MockValue; // Return all items if no filter value
    }

    const lowerCaseValue = filterValue.toLowerCase();

    return MockValue.filter(
      (item) => item.name.toLowerCase().includes(lowerCaseValue), // Filter by name property
    );
  }, [filterValue, MockValue]);

  const handleOnFocus = () => {
    setIsFocused(true);
  };

  const handleBlur = () => {
    setIsFocused(false);
  };

  const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFilterValue(e.target.value);
  };

  return (
    <div>
      <Input
        className=""
        onFocus={handleOnFocus}
        onBlur={handleBlur}
        onChange={handleOnChange}
        placeholder="search friends..."
      />
      {IsFocused && (
        <div className="mt-2 h-full w-full border-solid rounded-md border-2">
          {filteredItems.map((item, n) => (
            <>
              <div className="h-20 p-5 flex items-center hover:bg-gray-200">
                {item.name}
              </div>
              {n + 1 < filteredItems.length && (
                <Separator className="w-11/12 m-auto" />
              )}
            </>
          ))}
        </div>
      )}
    </div>
  );
}
