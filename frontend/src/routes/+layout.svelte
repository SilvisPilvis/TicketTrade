<script>
    import Cookies from "js-cookie";
    import axios from "$lib/axios";
    import { onMount } from "svelte";
    let loggedIn, isAdmin = false;

    onMount(async () => {
        try{
            // get all categories
            if(Cookies.get('token') != undefined){
                const config = {
                    headers: { Authorization: `Bearer ${Cookies.get('token')}` }
                };
                const resp = await axios.post('/admin', {}, config);
                console.log(resp.data);
                isAdmin = resp.data;
            }
        }catch (e){
            console.error("Error:", e.response);
            isAdmin = e.response.data;
            // failure = e.response.data.error;
        }
    })

    if(Cookies.get('token') != undefined){
        // window.location = 'http://localhost:5173/';
        loggedIn = true;
    }
    else{
        loggedIn = false;
    }

    function logout(){
        const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
        };
        axios.post('/logout', {}, config)
        .then(function (response) {
            console.log(response);
            Cookies.remove('token');
            loggedIn = false;
            window.location.replace("/login");
            // success = response.data
        })
        .catch(function (error) {
            Cookies.remove('token');
            window.location.replace("/login");
            console.log(error.response);
            // failed = error.response.data.error;
        });
    }
</script>

<nav class="flex row navbar">
    <div class="flex row">
        <a href="/" class="nav-button">Home</a>
        {#if loggedIn && isAdmin}
            <a href="/admin" class="nav-button">Admin</a>
        {/if}
    </div>
    {#if loggedIn != true}
        <div class="flex row">
            <a href="/login" class="nav-button">Login</a>
            <a href="/register" class="nav-button">Register</a>
        </div>   
    {/if}

    {#if loggedIn == true}    
        <button on:click={logout} class="nav-button">Logout</button>
    {/if}
</nav>

<style>
    .navbar{
        background-color: var(--button-fill);
        border-bottom-left-radius: 0.5rem;
        border-bottom-right-radius: 0.5rem;
        width: 100%;
        padding: 0.5rem 1rem;
    }

    .flex {
        display: flex;
        /* width: 100%; */
        flex-direction: row;
        /* justify-content: space-evenly; */
        justify-content: space-between;
        flex-wrap: nowrap;
    }
    .row {
        flex-direction: row;
    }
    nav {
        margin-bottom: 1rem;
        position: fixed;
        background-color: var(--bg);
    }
    .nav-button{
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0.5rem 1rem 0.5rem;
        border-radius: 0.4rem;
        background-color: var(--button-fill);
        color: var(--button-text);
        font-size: 1.25rem;
        font-weight: 800;
    }
    .nav-button:hover{
        filter: brightness(0.9);
        /* background-color: #B4ADEA; */
    }
    a {
        text-decoration: none;
        /* padding: 0 1rem 0 1rem; */
    }
    button{
        border: none;
        background-color: var(--bg);
        border-radius: 0.4rem;
        padding: 0.5rem 0 0.5rem;
    }
</style>

<slot/>