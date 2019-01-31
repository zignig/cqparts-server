<template>
        <div class="navbar-menu">
                <a class="button" @click="prevPage" :disabled="currentPage === 0">
                    <span class="icon"><i class="mdi mdi-chevron-double-left"></i></span>
                </a>
                <div class="">
                    <a class="button"
                        :key="model.name" 
                        v-for="model in pages"
                        v-on:click="load(model)">
                        <span class="icon" @click="pin(model)">
                            <i v-if="model.pinned" class="mdi mdi-pin"></i>
                            <i v-else="model.pinned" class="mdi mdi-rotate-45 mdi-pin"></i>
                        </span>
                        <span>{{ model.name }}</span>
                    </a>
                </div>
                    <a class="button" @click="nextPage" :is-visible="currentPage >= pageCount -1">
                        <span class="icon"><i class="mdi mdi-chevron-double-right"></i></span>
                    </a>
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
            default: 10,
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
            this.page[this.active]++;
        },
        prevPage: function(){
            this.page[this.active]--;
        },
        isPinned: function(obj){
            if (obj.pinned){
                return true
            }
        },
        isActive: function(obj){
            if (this.currenti != null ){
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
