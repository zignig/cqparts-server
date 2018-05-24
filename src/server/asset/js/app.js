source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

source.addEventListener('update',function(e){
    var data = JSON.parse(e.data);
    //console.log(data.Name);
    clear();
    load(data.Name);
    //d1.insertAdjacentHTML('beforeend',AlertBox(data));
});

var container, controls , obj;
var camera, scene, renderer, light;
var meshlist = [];
var raycaster ;
var mouse = new THREE.Vector2(), INTERSECTED ; 

init();
animate();

function clear(){
    scene.remove(obj);
}

function load(name){
	var loader = new THREE.GLTFLoader();
    loader.load( 'model/'+name+'/out.gltf', function ( gltf ) {
		scene.add( gltf.scene );
        obj = gltf.scene;
		//grab all the meshes for selection
        gltf.scene.traverse( function ( child ) {
            if ( child.isMesh ) {
                meshlist.push(child);
            }
            } );
        } );
    }

function init() {
	container = document.createElement( 'div' );
	document.body.appendChild( container );
	camera = new THREE.PerspectiveCamera( 30, window.innerWidth / window.innerHeight, 0.0001, 20 );
	camera.position.set( 0.2, 0.2, 0.2);
	controls = new THREE.OrbitControls( camera );
    //controls.autoRotate = true;
    controls.autoRotateSpeed = 2;
	controls.target.set(0,0,0);
	controls.update();
    scene = new THREE.Scene({antialias: true});
    scene.background =  new THREE.Color(255,255,255)
    // light 1
	light = new THREE.HemisphereLight( 0xbbbbff, 0x444422 );
	light.position.set( 0, 20, 0 );
    scene.add( light );
    // light 2
	light2 = new THREE.PointLight(0xf0f0f0,2,100);
	light2.position.set( 50,50,50);
	scene.add( light2 );
    // renderer
	renderer = new THREE.WebGLRenderer( { antialias: true } );
	renderer.setPixelRatio( window.devicePixelRatio );
	renderer.setSize( window.innerWidth, window.innerHeight );
	renderer.gammaOutput = true;
    // raycaster
    raycaster = new THREE.Raycaster();
    // add to doc and bind events
	container.appendChild( renderer.domElement );
	window.addEventListener( 'resize', onWindowResize, false );
	document.addEventListener('mousemove',onDocumentMouseMove,false);
	document.addEventListener('mousedown',onDocumentClick,false);
}

function onDocumentClick( event ) {
    if ( INTERSECTED ) {
        console.log(INTERSECTED);
        INTERSECTED.visible = false;
    };
}

function onDocumentMouseMove( event ) {
	event.preventDefault();
	mouse.x = ( event.clientX / window.innerWidth ) * 2 - 1;
	mouse.y = - ( event.clientY / window.innerHeight ) * 2 +1 ;
}

function onWindowResize() {
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize( window.innerWidth, window.innerHeight );
}

function animate() {
    requestAnimationFrame( animate );
    controls.update();
	render();
}

function render(){
	raycaster.setFromCamera(mouse,camera);
	var intersects = raycaster.intersectObjects(meshlist);
	if (intersects.length > 0){
        //console.log(intersects);
		if ( INTERSECTED != intersects[0].object){
            if (INTERSECTED) INTERSECTED.material.emissive.setHex(INTERSECTED.currentHEX);
			INTERSECTED = intersects[0].object;
			INTERSECTED.currentHex = INTERSECTED.material.emissive.getHex();
			INTERSECTED.material.emissive.setHex(0x333313);
		}
	} else {
        if ( INTERSECTED ) INTERSECTED.material.emissive.setHex(INTERSECTED.currentHex);
        INTERSECTED = null;
	}
    renderer.render( scene, camera );
}
