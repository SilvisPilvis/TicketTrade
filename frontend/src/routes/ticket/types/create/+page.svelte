<script>
    import axios from "$lib/axios"
    let success, fail, priceFail, typeName, typePrice;
    function createTicketType(){
        // console.log(typeName);
        console.log(typePrice);
        if(typePrice <= 0 || typePrice === undefined){
            priceFail = "Ticket Type Price must be greater than 0.";
        }else{
            priceFail = "";
        }
        // if(typePrice === undefined){
        //     console.log("test");
        //     fail = "Ticket Type Price must be greater than 0.";
        // }
        if(typeName == "" || typeName === undefined){
            fail = "Ticket Type Name can't be empty.";
        }else{
            fail = "";
        }
        if(typePrice > 0 && typePrice != undefined && typeName != "" && typeName != undefined){
            axios.post('/ticket/type/create', {
                typeName: typeName,
                typePrice: typePrice
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
    }
</script>

<main class="flex col cen">
    <div class="margin-t flex col">
        <h1>Create Ticket Type</h1>
        <input type="text" name="" id="" bind:value={typeName} placeholder="Type name">
        <input type="number" name="" id="" min="0.01" step="0.01" bind:value={typePrice} placeholder="Type Price">
        <button on:click={createTicketType}>Create Ticket Type</button>
        {#if fail != "" && fail != undefined}
            <p class="error">{fail}</p>
        {/if}
        {#if priceFail != "" && priceFail != undefined}
            <p class="error">{priceFail}</p>
        {/if}
        {#if success != "" && success != undefined}
            <p class="success">{success}</p>
        {/if}
    </div>
</main>

<style>
    main{
        height: 100%;
    }
    .margin-t{
        margin-top: 5rem;
    }
    input, button{
        margin: 1rem;
        padding: 0.6rem;
        border: none;
        border-radius: 0.4rem;
        background-color: var(--fg);
    }
    button{
        background-color: var(--button-fill);
    }
</style>