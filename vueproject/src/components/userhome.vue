<template>
    <diV>
        <div style="height: 50px"></div>
        <el-col :offset="4">
        <el-container>
            <el-aside width="450px">
                <div class="asidebox">
                    <div style="height: 100px;">
                        <i class="el-icon-user-solid" style="font-size:50px; line-height: 150px;"></i>
                    </div>
                    <div>
                        <p>{{ username }}</p>
                    </div>
                </div>
            </el-aside>
            <el-main id='usermain'></el-main>
        </el-container>
        </el-col>
    </diV>
</template>


<script>
export default{
    name:'userhome',
    props:['username'],
}
</script>

<style>
#usermain {
    background-color: rgb(118, 223, 195);
    padding: 0%;
}
.el-aside {
    background-color: aqua;
    padding: 0%;
}
.asidebox {
    text-align: center;
    
}

</style>