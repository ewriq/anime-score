<script>
  import { writable } from "svelte/store";
  import Toast from "../Toast.svelte";
  import axios from "axios";

  export let voter;
  export let anime;
  export let score;

  const toastVisible = writable(false);
  let toastMessage;

  function showToast(message) {
    toastMessage = message;
    toastVisible.set(true);

    setTimeout(() => {
      toastVisible.set(false);
    }, 3000);
  }
  async function del() {
    try {
      const response = await axios.post("http://localhost:3000/api/vote/del", {
        voter: voter,
        anime: anime,
        score: score,
      });
      if (response.data.status == 502) {
        toastMessage = "Sunucu tarafıyla ilgili sorun oluştu";
        showToast(toastMessage);
      } else {
        toastMessage = "Başarılı bir şekilde oyunuz kaldırıldı";
        showToast(toastMessage);
        setTimeout(() => {
          window.location.href = "/";
        }, 4000);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  }
</script>

<div>
  <button on:click={del}> 🚮 </button>
  <Toast message={toastMessage} bind:visible={$toastVisible} />
</div>
