<script>
    import axios from '../../services/axios';
    let username, password, password_repeat, usernameError, passwordError, repeatError;
    
    function login(){
        axios.post('/login', {
            username: username,
            password: password
        })
        .then(function (response) {
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
    }
    </script>
    
    <main>
    <input type="text" name="username" id="" bind:value={username} placeholder="username">
    <input type="password" name="password" id="" bind:value={password} placeholder="password">
    <button on:click={login}>Login</button>
    </main>
    
    <style>
    
    </style>