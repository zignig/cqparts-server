<template>
    <b-notification class="middle" :active.sync="isActive" >
        <div v-bind:model="issueItem">{{ issueItem }}</div>
    </b-notification>
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
