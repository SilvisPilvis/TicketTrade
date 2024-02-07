<script>
    import Cookies from "js-cookie";
    import axios from "$lib/axios";
    function logout(){
        const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
        };
        axios.post('/logout', {}, config)
        .then(function (response) {
            console.log(response);
            Cookies.remove('token');
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

<nav class="flex row">
    <a href="/" class="nav-button">Home</a>
    <a href="/login" class="nav-button">Login</a>
    <a href="/register" class="nav-button">Register</a>
    <button on:click={logout} class="nav-button">Logout</button>
</nav>

<style>
    .flex {
        display: flex;
        width: 100%;
        flex-direction: row;
        /* justify-content: space-evenly; */
        justify-content: space-around;
        /* flex-wrap: wrap; */
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
        padding: 0.5rem 0 0.5rem;
        border-radius: 0.4rem;
        background-color: #B4ADEA;
        color: #4F0147;
        font-size: 1.25rem;
        font-weight: 800;
    }
    .nav-button:hover{
        filter: brightness(0.9);
        /* background-color: #B4ADEA; */
    }
    a {
        text-decoration: none;
    }
    button{
        border: none;
        background-color: var(--bg);
        border-radius: 0.4rem;
        padding: 0.5rem 0 0.5rem;
    }
</style>

<slot/>