'use strict';

document.querySelector("#register").addEventListener("submit", async (e) => {
    e.preventDefault();

    var test = new FormData(e.target);
    console.log(test.get("username"));
    console.log(test.get("password"));

    const username = new FormData(e.target).get("username");
    console.log(username);

    const json = await fetchJson("register", {
        method: "POST",
        credentioals: "include",
        body: new FormData(e.target)
    });

    if (json.status === "200") {
        if (navigator.credentials) {
            var cred = new PasswordCredential({
                id: test.get("username"),
                password: test.get("password")
            });
            navigator.credentials.store(cred).then(() => {
                console.log("You are registered");
            });
        }
    }

    console.log(json);
});

document.querySelector("#login").addEventListener("click", async () => {
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
