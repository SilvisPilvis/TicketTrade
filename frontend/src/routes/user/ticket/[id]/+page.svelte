<script>
    import { onMount } from "svelte";
    import axios from "$lib/axios";
    import html2PDF from 'jspdf-html2canvas';
    let success, failed;
    // import html2PDF from 'jspdf-html2canvas';
    // import html2PDF from "html-pdf-adaptive";

    
    export let data;
    let res;
    onMount(async () => {
        try{
            const response = await axios.get(`/bought/ticket/${data.data}`);
            res = response.data;
            console.log(res);
        }catch (e){
            console.error("Error:", e);
            failed = e.response.data.error;
        }
    })
    function getPDF(){
        const el = document.querySelector(".ticket")
        html2PDF(el, {
        jsPDF: {
          format: 'a4',
        },
        imageType: 'image/png',
        output: './ticket.pdf'
      });
    }

</script>

<main class="flex col cen">
    {#if res != "" && res != undefined}
    <div class="margin-t flex col cen">
        <div class="ticket">
            <div class="flex row">
                <img class="square-img" src={res.TicketImage} alt="">
                <!-- <img class="square-img" src={image} alt=""> -->
                <div class="ticket-data">
                    <div class="flex row evenly">
                        <p>{res.UserName}</p>
                        <div> </div>
                        <p class="title">{res.EventName}</p>    
                    </div>
                    <p class="evenly">Seat: {res.TicketSeat}</p>
                    <p class="evenly">{res.TicketDate}</p>
                    <p class="evenly">{res.TicketLocation}</p>
                </div>
                <div class="qr-box flex cen">
                    <img class="qr" src={res.QrPath.replace("http://localhost:8000/", "")} alt="">
                </div>
            </div>
        </div>
        <button on:click={getPDF}>Download As PDF</button>
    </div>
    {/if}
</main>

<style>
    .margin-t{
        margin-top: 5rem;
        width: 100%;
    }
    @-webkit-keyframes rainbow {
      from {
        box-shadow:0 0 8px 8px var(--fg);
      }
      25%{
        box-shadow:0 0 8px 7px var(--fg);
      }
      50%{
        box-shadow:0 0 8px 6px var(--fg);
      }
      75%{
        box-shadow:0 0 8px 7px var(--fg);
      }
      to {
        box-shadow:0 0 8px 8px var(--fg);
      }
    }
    .ticket{
        margin-left: 1rem;
        max-width: 57rem;
        width: 57rem;
        /* outline: 2px solid var(--fg); */
        border-radius: 0.25rem;
        background-color: var(--fg);
/*        box-shadow: 0px 0px 8px 7px blue;*/
        -webkit-animation:rainbow 1s linear infinite;
    }
    .ticket-data{
        width: 43%;
    }
    .title{
        font-weight: 700;
    }
    .evenly{
        font-size: 16pt;
        display: flex;
        margin-left: 2rem;
        height: 20%;
        align-items: center;
    }
    .square-img{
        object-fit: cover;
        border-right: 2px solid var(--fg);
        max-height: 20rem;
    }
    .qr-box{
        aspect-ratio: 1 / 1;
/*        height: 100%;*/
        border-left: 1px dashed var(--fg);
        border-width:3px;
    }
    button{
        margin-top: 3rem;
        margin-left: 1rem;
        outline: none;
        border: none;
        padding: 0.75rem;
        background-color: var(--button-fill);
        /* color: var(--fg); */
        border-radius: 0.25rem;
    }
</style>