<script>
    import { onMount } from "svelte";
    import axios from "$lib/axios"
    export let data;
    let res;
onMount(async () => {
    try{
        const response = await axios.get(`/${data.data}/tickets`);
        res = response.data;
        console.log(res);
    }catch (e){
        console.error("Error:", e);
        failed = e.response.data.error;
    }
})
</script>

<main>
    {#if res != "" && res != undefined}
        {#each res as ticket}
            <div>
                <p>{ticket.EventName}</p>
                <p>{ticket.UserName}</p>
                <p>{ticket.TicketBoughtAt}</p>
                <p>{ticket.TicketLocation}</p>
                <p>{ticket.TicketDate}</p>
                <p>{ticket.TicketSeat}</p>
            </div>
        {/each}
        <!-- <div>
            <img src={"http://localhost:8000/"+res.QrPath} alt="">
            <p>{res.UserName}</p>
            <p>{res.EventName}</p>
            <p>{res.TicketDate}</p>
            <p>{res.TicketLocation}</p>
            <p>{res.TicketSeat}</p>
        </div> -->
    {/if}
</main>