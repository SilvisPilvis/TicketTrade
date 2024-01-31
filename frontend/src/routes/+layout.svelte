<nav class="flex row">
    <a href="/"><div class="nav-button">Home</div></a>
    <a href="/login"><div class="nav-button">Login</div></a>
    <a href="/register"><div class="nav-button">Register</div></a>
</nav>

<style>
    .flex {
        display: flex;
        flex-wrap: wrap;
    }
    .row {
        flex-direction: row;
    }
    nav {
        position: absolute;
    }
    .nav-button{
        width: calc(100vw/3);
        padding: 0.3rem 0 0.3rem;
        background-color: rgba(28, 92, 183, 0.689);
    }
    a {
        text-decoration: none;
        color: black;;
    }
</style>

<slot/>