<script>
  import axios from "axios";
  import { writable } from "svelte/store";
  import Cookies from "js-cookie";
  import Toast from "../Toast.svelte";

  export let anime;

  let user = Cookies.get("token");
  let rating = 0;
  let selected = 0;

  const toastVisible = writable(false);
  let toastMessage;

  function showToast(message) {
    toastMessage = message;
    toastVisible.set(true);

    setTimeout(() => {
      toastVisible.set(false);
    }, 3000);
  }

  async function sendRatingToAPI() {
    try {
      const response = await axios.post("http://localhost:3000/api/vote/add", {
        score: selected,
        anime: anime,
        voter: user,
      });
      if (response.data.status == 502) {
        toastMessage = "Daha önce oy kullanmış olabilirsiniz";
        showToast(toastMessage);
      } else {
        toastMessage = "Başarılı bir şekilde oy kullandınız";
        showToast(toastMessage);
        setTimeout(() => {
          window.location.href = "/anime/" + anime;
        }, 4000);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  }

  function handleStarClick(num) {
    selected = num;
  }

  function handleSubmit() {
    if (selected > 0) {
      rating = selected;
      sendRatingToAPI();
    } else {
      console.error("Bir rating seçmelisiniz.");
    }
  }

  const Star = ({ filled }) => `
        <svg class="w-6 h-6 ${filled ? "text-yellow-400" : "text-gray-300"} mx-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 22 20">
            <path d="M20.924 7.625a1.523 1.523 0 0 0-1.238-1.044l-5.051-.734-2.259-4.577a1.534 1.534 0 0 0-2.752 0L7.365 5.847l-5.051.734A1.535 1.535 0 0 0 1.463 9.2l3.656 3.563-.863 5.031a1.532 1.532 0 0 0 2.226 1.616L11 17.033l4.518 2.375a1.534 1.534 0 0 0 2.226-1.617l-.863-5.03L20.537 9.2a1.523 1.523 0 0 0 .387-1.575Z"/>
        </svg>
    `;
</script>

<div class="flex items-center">
  {#each [1, 2, 3, 4, 5] as num}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div on:click={() => handleStarClick(num)} class="cursor-pointer">
      {@html Star({ filled: selected >= num })}
    </div>
  {/each}
  {#if selected > 0}
    <button
      class="ml-4 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700 transition-colors"
      on:click={handleSubmit}
    >
      Gönder
    </button>
  {/if}
  <!-- Bind the toastVisible writable store to the visible prop of Toast component -->
  <Toast message={toastMessage} bind:visible={$toastVisible} />
</div>
