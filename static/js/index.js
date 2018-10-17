
document.querySelector("#register").addEventListener("submit", async (e) => {
    e.preventDefault();

    const username = new FormData(e.target).get("username");
    console.log(username);

    const responce = await fetch("register", {
        method: "GET",
    });

    const json = await responce.json();

    console.log(json);
});
