<script lang="ts">
    export let name: string;

</script>

<main>
    <h1>Hello {name}!</h1>
    <p>
        Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn
        how to build Svelte apps.
    </p>
    <a href="https://id.twitch.tv/oauth2/authorize?client_id=lc9rwae6zjagp8efydksaow47sgo2v&redirect_uri=http://localhost:5000/twitch-auth&scope=openid&response_type=code"><button>Let me login</button></a>
</main>

<style lang="scss">
    h1 {
        color: #ff3e00;
        text-transform: uppercase;
        font-size: 4em;
        font-weight: 100;
    }
    
    main {
        text-align: center;
        padding: 1em;
        max-width: 240px;
        margin: 0 auto;
    }
    
    @media (min-width: 640px) {
        main {
            max-width: none;
        }
    }
</style>
