document.querySelector("#register").addEventListener("submit", async (e) => {
    e.preventDefault();

    const username = new FormData(e.target).get("username");
    console.log(username);

    //const responce = await fetch("register", {
    //    method: "GET",
    //});

    //const json = await responce.json();
    //
    const json = await fetchJson("register", {
        method: "GET"
    });

    console.log(json);
});

async function fetchJson(url, option) {
    
    const res = await fetch(url, option);

    return res.json();
}
