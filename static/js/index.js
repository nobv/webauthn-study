'use strict';

document.querySelector("#register").addEventListener("submit", async (e) => {
    e.preventDefault();

    const username = new FormData(e.target).get("username");
    console.log(username);

    const json = await fetchJson("register", {
        method: "GET"
    });

    console.log(json);
});

document.querySelector("#test").addEventListener("click", async () => {
    navigator.credentials.get({'password': true}).then(credential => {
        if(!credential){
            throw new Error("No Credential returned");
        }

        let credentials = {
            'username': credential.id,
            'password': credential.password
        }
    });
});

async function fetchJson(url, option) {
    
    const res = await fetch(url, option);

    return res.json();
}
