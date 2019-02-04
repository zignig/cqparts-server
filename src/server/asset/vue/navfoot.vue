<template>
    <nav class="navbar is-fixed-bottom">
        <div class="navbar-brand">
            <a class="button" id="autorotate" onclick="ActivateAutorotate()">Autorotate</a>
            <a class="button" id="render" onclick="PostRender()">PostRender</a>
        </div>
    </nav>
</template>

<script>
export default {
    data() {
        return {
            isActive: false
        }
    },
    props: {
        issueItem: {
            type: String,
        },
    },
    methods:{
    }
}
</script>

<style>
</style>
