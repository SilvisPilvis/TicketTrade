<script>
    import axios from "$lib/axios";
    let password, email, fail, success;

    function resetPassword(){
        if(password == "" || password == undefined){
            fail = "New password can't be empty.";
        }else{
            fail = "";
        }
        if(email == "" || email == undefined){
            fail = "Email can't be empty."
        }else{
            fail = "";
        }

        axios.put('/password/reset', {
            email: email,
            password: password,
        })
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            fail = error.response.data.error;
            success = "";
        });
    }
</script>

<main>
    <input type="email" bind:value={email}>
    <input type="password" name="" id="" bind:value={password}>
    <button on:click={resetPassword}>Reset Password</button>
    {#if fail != undefined && fail != ""}
        <p class="error">{fail}</p>
    {/if}
    {#if success != undefined && success != ""}
        <p class="success">{success}</p>
    {/if}
</main>