<script>
    import axios from '$lib/axios';
    let username, email, password, passwordRepeat, usernameError = false, emailError = false, passwordError = false, repeatError = false, success, fail;
    
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
            console.log(response.data);
            success = response.data;
            fail = "";
        })
        .catch(function (error) {
            console.log(error);
            fail = error.response.data.error;
            success = "";
        });
    }
    </script>
    
    <main class="flex col cen">
        <div class="margin-t flex col">
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
            {#if fail != "" && fail != undefined}
                <p class="error">{fail}</p>
            {/if}
            {#if success != "" && success != undefined}
                <p class="success">{success}</p>
            {/if}
            <button on:click={register}>Register</button>
        </div>
    </main>
    
    <style>
    .margin-t{
        margin-top: 5rem;   
    }
    .error {
        color: red;
        font-size: 1.3rem;
        height: min-content;
        margin: 0;
    }
    p.error, p.success{
        max-width: 14rem;
    }
    input {
        display: flex;
        padding: 1rem;
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: var(--bg);
        font-size: 12pt;
        margin: 1rem;
    }
    button {
        padding: 0.7rem 6rem 0.7rem 6rem;
        border: none;
        outline: none;
    }
    </style>