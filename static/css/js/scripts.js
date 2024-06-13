document.addEventListener("DOMContentLoaded", function () {
    const addUserForm = document.getElementById("add-user-form");
    const userList = document.getElementById("users");

    addUserForm.addEventListener("submit", function (event) {
        event.preventDefault(); // Prevent the form from submitting

        const formData = new FormData(addUserForm);
        const userData = {
            name: formData.get("name"),
            email: formData.get("email")
        };

        fetch("/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(userData)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error("Failed to add user");
            }
            return response.json();
        })
        .then(data => {
            // Clear form inputs
            addUserForm.reset();
            // Add new user to the list
            renderUser(data);
        })
        .catch(error => {
            console.error("Error:", error);
            alert("Failed to add user. Please try again.");
        });
    });

    // Function to render a user item in the list
    function renderUser(user) {
        const userItem = document.createElement("li");
        userItem.textContent = `${user.name} (${user.email})`;
        userList.appendChild(userItem);
    }

    // Fetch all users on page load
    fetch("/users")
    .then(response => response.json())
    .then(users => {
        users.forEach(user => renderUser(user));
    })
    .catch(error => {
        console.error("Error fetching users:", error);
        alert("Failed to fetch users. Please refresh the page.");
    });
});
