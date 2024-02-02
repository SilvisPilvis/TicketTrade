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
            })
            .catch(function (error) {
                console.log(error);
                fail = error.response.data.error;
            });
    }
</script>

<main>
    <input type="text" name="typeName" id="" bind:value={typeName}>
    <input type="number" name="typePrice" id="" min="0.01" step="0.01" bind:value={typePrice}>
    {#if failure != "" && failure != undefined}
        <p class="error">{failure}</p>
    {/if}
    {#if success != "" && success != undefined}
        <p class="error">{success}</p>
    {/if}
    <button on:click={editTicketType}>Edit Ticket Type</button>
</main>