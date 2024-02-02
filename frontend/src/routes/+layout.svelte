<nav class="flex row">
    <a href="/" class="nav-button">Home</a>
    <a href="/login" class="nav-button">Login</a>
    <a href="/register" class="nav-button">Register</a>
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
    }
    .nav-button{
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0.5rem 0 0.5rem;
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
</style>

<slot/>