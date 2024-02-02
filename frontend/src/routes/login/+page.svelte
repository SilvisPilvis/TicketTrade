<script>
    import axios from '$lib/axios';
    import Cookies from 'js-cookie';
    let username, password, fail, usernameError = false, passwordError = false;

    if(Cookies.get('token') != undefined){
        window.location = 'http://localhost:5173/';
    }
    
    function login(){
        if(username == "" || username == undefined){
            usernameError = true;
        }else{
            usernameError = false;
        }
        if(password == "" || password == undefined){
            passwordError = true;
            // return;
        }else{
            passwordError = false;
        }
        if(password == undefined || password == "" || username == undefined || username == ""){
            return;
        }
        fail = "";
        axios.post('/login', {
            username: username,
            password: password
        })
        .then(function (response) {
            Cookies.set("token", response.data);
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
            fail = error.message;
        });
    }
</script>
    
<main class="flex cen">
    <div class="flex cen col">
        <input type="text" name="username" id="" bind:value={username} placeholder="username">
        {#if usernameError != false}
            <p class="error">Username can't be empty.</p>
        {/if}
        <input type="password" name="password" id="" bind:value={password} placeholder="password">
        {#if passwordError != false}
            <p class="error">Password can't be empty.</p>
        {/if}
        {#if fail != "" && fail != undefined}
            <p class="error">{fail}</p>
        {/if}
        <button on:click={login}>Login</button>
    </div>
</main>

<style>
    main{
        height: 100vh;
        width: 100vw;
    }
    .flex {
        display: flex;
        flex-wrap: wrap;
    }
    .cen {
        justify-content: center;
        align-items: center;
    }
    .col {
        flex-direction: column;
    }
    .error {
        color: red;
        font-size: 1.3rem;
        height: min-content;
        margin: 0;
        /* background-color: ; */
        /* text-shadow: 10px, 10px, 1px, rgb(97, 10, 10); */
    }
    input {
        padding: 1rem;
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: gray;
        font-size: 12pt;
        margin: 1rem;
    }
    button {
        padding: 0.7rem 6rem 0.7rem 6rem;
        border: none;
        outline: none;
    }
</style>