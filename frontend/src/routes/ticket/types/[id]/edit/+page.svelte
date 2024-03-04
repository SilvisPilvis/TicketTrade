<script>
    import axios from "$lib/axios";
    import { onMount  } from "svelte";
    export let data;
    let typeName, typePrice, result, success, failure;
    
    onMount(async () => {
        try{
            const res = await axios.get(`/ticket/type/${data.data}`);
            // const res = await axios.get('/events');
            result = res.data;
            typeName = result.typeName;
            typePrice = result.typePrice;
        }catch (e){
            console.error("Error:", e.response.data.error);
            failure = e.response.data.error;
        }
    })

    function editTicketType(){
        if(typeName == "" || typeName == undefined){
            failure = "Ticket Type Name can't be empty.";
            return;
        }else{
            failure = "";
        }
        if(typePrice <= 0){
            failure = "Ticket Type Price must be greater than 0.";
            return;
        }else{
            failure = "";
        }
        axios.put(`/ticket/type/${data.data}`, {
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
</script>

<main class="flex col cen">
    <div class="margin-t flex col">
        <h1>Edit Ticket Type</h1>
        <input type="text" name="typeName" id="" bind:value={typeName}>
        <input type="number" name="typePrice" id="" min="0.01" step="0.01" bind:value={typePrice}>
        {#if failure != "" && failure != undefined}
            <p class="error">{failure}</p>
        {/if}
        {#if success != "" && success != undefined}
            <p class="success">{success}</p>
        {/if}
        <button on:click={editTicketType}>Edit Ticket Type</button>
    </div>
</main>

<style>
    main{
        height: 100%;
    }
    .margin-t{
        margin-top: 5rem;
    }
    input{
        background-color: var(--fg);
    }
    button{
        background-color: var(--button-fill);
    }
    button, input{
        border: none;
        margin: 0.5rem;
        padding: 0.7rem;
        border-radius: 0.4rem;
    }
</style>