<script>
    import axios from "$lib/axios"
    let success, fail, typeName, typePrice;
    function createTicketType(){
        if(typeName == "" || typeName == undefined){
            fail = "Ticket Type Name can't be empty.";
        }else{
            fail = "";
        }
        if(typePrice <= 0){
            fail = "Ticket Type Price must be greater than 0.";
        }else{
            fail = "";
        }
        if(typePrice > 0 && typeName != "" && typeName != undefined){
            axios.post('/ticket/type/create', {
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
    }
</script>

<main>
    <input type="text" name="" id="" bind:value={typeName}>
    <input type="number" name="" id="" min="0.01" step="0.01" bind:value={typePrice}>
    <button on:click={createTicketType}>Create Ticket Type</button>
    {#if fail != "" && fail != undefined}
        <p class="error">{fail}</p>
    {/if}
    {#if success != "" && success != undefined}
        <p class="error">{success}</p>
    {/if}
</main>