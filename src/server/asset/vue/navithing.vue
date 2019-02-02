<template>
        <div class="navbar-menu">
                <div class="buttons has-addons">
                <a class="button" @click="prevPage" v-show="pageCount > 1">
                    <span class="icon"><i class="mdi mdi-chevron-double-left"></i></span>
                </a>
                <a class="button"
                    :key="model.name" 
                    v-for="model in pages"
                    v-on:click="load(model)">
                    <span>{{ model.name }}</span>
                </a>
                <a class="button" @click="nextPage" v-show="pageCount > 1" :is-visible="currentPage >= pageCount -1">
                    <span class="icon"><i class="mdi mdi-chevron-double-right"></i></span>
                </a>
                </div>
        </div>
</template>

<script>
export default {
    data(){
        return {
            page: { 'All':0,'Pinned':0,'Fix':0},
            visible: true,
            active: 'All',
            sections: ['Pinned','All','Fix']
        }
    },
    props: {
        current: "",
        modelList: {
            type: Array,
        },
        size:{
            type: Number,
            required: false,
            default: 8,
        }
    },
    methods:{
        sectionActive: function(name){
            if (this.active == name){
                return "is-active"
            }
        },
        selectSection: function(name){
            if (this.visible){
                this.active = name
            }
        },
        viz: function(){
            this.visible = !this.visible;
        },
        nextPage: function(){
            if ( this.page[this.active] < this.pageCount-1){
                this.page[this.active]++;
            }
        },
        prevPage: function(){
            if ( this.page[this.active] > 0 ){
                this.page[this.active]--;
            }
        },
        isPinned: function(obj){
            if (obj.pinned){
                return true
            }
        },
        isActive: function(obj){
            if (this.current != null ){
                if (this.current.name  == obj.name){
                    return "green"
                }
                else {
                    return "black"
                }
            }
        },
        pin: function(obj){
            obj.pinned = !obj.pinned
            EventBus.$emit('pin',obj.name);
        },
        load: function(obj){
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        },
        pinned : function(model){
            return this.modelList.filter(function(model){
                return model.pinned === true
            })
        },
        currentData: function(){
            let data = []
            if ( this.active  == 'Pinned' ){
                data = this.pinned()
            }
            if ( this.active  == 'All' ){
                data = this.modelList
            }
            if ( this.active  == 'Fix' ){
                data = this.modelList
            }
            return data
        },
        paginatedData: function(){
            const start = this.page[this.active]* this.size,
                end = start + this.size;
            data = this.currentData()
            return data.slice(start, end);
        }
    },
    computed:{
        currentPage(){
            return this.page[this.active]
        },
        pageCount(){
            let l = this.currentData().length,
                s = this.size;
            return Math.ceil(l/s);
        },
        pages() {
            return this.paginatedData();
        }
    }
};

</script>

<style>
.nogap {
    padding-top: 0.0 !important;
    padding-right: 0.0 !important;
    padding-left: 0.0 !important;
    padding-bottom: 0.0 !important;
}
</style>
