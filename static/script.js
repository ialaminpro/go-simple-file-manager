document.addEventListener("DOMContentLoaded", () => {
    const copyButtons = document.querySelectorAll(".copy-btn");

    copyButtons.forEach(button => {
        button.addEventListener("click", () => {
            const url = button.getAttribute("data-url");
            navigator.clipboard.writeText(`${window.location.origin}${url}`).then(() => {
                alert("URL copied to clipboard!");
            }).catch(err => {
                alert("Failed to copy URL: " + err);
            });
        });
    });

    // Handle Delete File
    document.querySelectorAll(".delete-btn").forEach((button) => {
        button.addEventListener("click", async () => {
            const fileName = button.getAttribute("data-file");
            const confirmed = confirm(`Are you sure you want to delete "${fileName}"?`);
            if (!confirmed) return;

            try {
                const response = await fetch(`/delete`, {
                    method: "DELETE",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ file: fileName }),
                });
                if (response.ok) {
                    alert("File deleted successfully.");
                    location.reload();
                } else {
                    alert("Failed to delete file.");
                }
            } catch (error) {
                alert("An error occurred: " + error.message);
            }
        });
    });

    // Handle Rename File
    document.querySelectorAll(".rename-btn").forEach((button) => {
        button.addEventListener("click", async () => {
            const fileName = button.getAttribute("data-file");
            const newName = prompt(`Enter a new name for "${fileName}":`);
            if (!newName) return;

            try {
                const response = await fetch(`/rename`, {
                    method: "PUT",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ oldName: fileName, newName }),
                });
                if (response.ok) {
                    alert("File renamed successfully.");
                    location.reload();
                } else {
                    alert("Failed to rename file.");
                }
            } catch (error) {
                alert("An error occurred: " + error.message);
            }
        });
    });
});
