<script>
    import axios from '$lib/axios';
    let username, email, password, passwordRepeat, usernameError = false, emailError = false, passwordError = false, repeatError = false;
    
    function register(){
        if(username == "" || username == undefined){
            usernameError = true;
        }else{
            usernameError = false;
        }
        if(email == "" || email == undefined){
            emailError = true;
        }else{
            emailError = false;
        }
        if(password == "" || password == undefined){
            passwordError = true;
            // return;
        }else{
            passwordError = false;
        }
        if(password != passwordRepeat){
            repeatError = true;
        }else{
            repeatError = false;
        }
        if(password == undefined || password == "" || email == undefined || email == "" || username == undefined || username == "" || passwordRepeat == undefined || passwordRepeat == ""){
            return;
        }
        axios.post('/register', {
            username: username,
            email: email,
            password: password
        })
        .then(function (response) {
            console.log(response.data.success);
        })
        .catch(function (error) {
            console.log(error);
        });
    }
    </script>
    
    <main class="flex col cen">
        <input type="text" name="username" id="" bind:value={username} placeholder="username">
        {#if usernameError != false}
            <p class="error">Username can't be empty.</p>
        {/if}
        <input type="email" name="email" id="" bind:value={email} placeholder="email">
        {#if emailError != false}
            <p class="error">Email can't be empty.</p>
        {/if}
        <input type="password" name="password" id="" bind:value={password} placeholder="password">
        {#if passwordError != false}
            <p class="error">Password can't be empty.</p>
        {/if}
        <input type="password" id="" bind:value={passwordRepeat} placeholder="password">
        {#if repeatError != false}
            <p class="error">Passwords must match.</p>
        {/if}
        <button on:click={register}>Register</button>
    </main>
    
    <style>
        * , *::after, *::before{
            box-sizing: border-box;
            min-width: 0;
            padding: 0;
            margin: 0;
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