import { Loader2, UserRound } from "lucide-react";
import { Button } from "./ui/button";
import { Textarea } from "./ui/textarea";
import { ReactEventHandler, useState } from "react";
import apiClient from "@/lib/apiClient";

export const CreatedPost = () => {
  const [value, setValue] = useState("");
  const [isLoading, setIsLoading] = useState(false); // State for button loading indicator
  const [error, setError] = useState(""); // State for error handling

  const handlePostSubmit: ReactEventHandler<HTMLButtonElement> = async (
    event,
  ) => {
    event.preventDefault(); // Prevent default form submission behavior

    if (!value) {
      setError("Please enter your post content.");
      return; // Exit function early if no content entered
    }

    setIsLoading(true); // Set loading indicator for button
    setError(""); // Clear any previous errors

    try {
      const response = await apiClient.post("/post", {
        content: value,
        visibility_type_id: 1,
        user_uuid: "549e9b06-4792-4b45-8ce8-63b7b93be7a7",
      }); // Send POST request
      console.log("Post created successfully:", response.data); // Log response for debugging

      // Handle successful post creation (e.g., clear form, show success message)
      setValue("");
    } catch (error) {
      console.error("Error creating post:", error); // Log error for debugging
      setError("An error occurred. Please try again later."); // Set error message
    } finally {
      setIsLoading(false); // Reset loading indicator after request completes
    }
  };
  return (
    <div className="my-10 ">
      <div className="flex flex-col items-center">
        <div className="flex items-center gap-4 w-full">
          <div className="w-8 h-8 rounded-full bg-gray-800">
            <UserRound size={32} color="white" />
          </div>
          <Textarea
            placeholder="Type your stories!"
            value={value}
            onChange={(event) => setValue(event.target.value)}
          />
        </div>
        <Button
          onClick={handlePostSubmit}
          className="w-4/5 mt-2"
          disabled={value === ""}
        >
          {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          Post
        </Button>
        {error && <p className="text-destructive">{error}</p>}
      </div>
    </div>
  );
};
