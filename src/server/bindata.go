// Code generated by go-bindata.
// sources:
// asset/css/extra.css
// asset/css/semantic.min.css
// asset/css/themes/default/assets/fonts/icons.eot
// asset/css/themes/default/assets/fonts/icons.svg
// asset/css/themes/default/assets/fonts/icons.ttf
// asset/css/themes/default/assets/fonts/icons.woff
// asset/css/themes/default/assets/fonts/icons.woff2
// asset/html/buttons.tmpl
// asset/html/dev.tmpl
// asset/html/index.tmpl
// asset/html/viewer.tmpl
// asset/js/GLTFLoader.js
// asset/js/OrbitControls.js
// asset/js/app.js
// asset/js/axios.min.js
// asset/js/components.js
// asset/js/cqpartsViewer.js
// asset/js/saver.js
// asset/js/semantic-ui-vue.min.js
// asset/js/three.min.js
// asset/js/vue.js
// asset/js/vuex.js
// asset/vue/App.vue
// asset/vue/card.vue
// asset/vue/model.vue
// asset/vue/notif.vue
// asset/vue/steps.vue
// DO NOT EDIT!

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assetCssExtraCss reads file data from disk. It returns an error on failure.
func assetCssExtraCss() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/extra.css"
	name := "asset/css/extra.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssSemanticMinCss reads file data from disk. It returns an error on failure.
func assetCssSemanticMinCss() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/semantic.min.css"
	name := "asset/css/semantic.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssThemesDefaultAssetsFontsIconsEot reads file data from disk. It returns an error on failure.
func assetCssThemesDefaultAssetsFontsIconsEot() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/themes/default/assets/fonts/icons.eot"
	name := "asset/css/themes/default/assets/fonts/icons.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssThemesDefaultAssetsFontsIconsSvg reads file data from disk. It returns an error on failure.
func assetCssThemesDefaultAssetsFontsIconsSvg() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/themes/default/assets/fonts/icons.svg"
	name := "asset/css/themes/default/assets/fonts/icons.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssThemesDefaultAssetsFontsIconsTtf reads file data from disk. It returns an error on failure.
func assetCssThemesDefaultAssetsFontsIconsTtf() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/themes/default/assets/fonts/icons.ttf"
	name := "asset/css/themes/default/assets/fonts/icons.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssThemesDefaultAssetsFontsIconsWoff reads file data from disk. It returns an error on failure.
func assetCssThemesDefaultAssetsFontsIconsWoff() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/themes/default/assets/fonts/icons.woff"
	name := "asset/css/themes/default/assets/fonts/icons.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetCssThemesDefaultAssetsFontsIconsWoff2 reads file data from disk. It returns an error on failure.
func assetCssThemesDefaultAssetsFontsIconsWoff2() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/css/themes/default/assets/fonts/icons.woff2"
	name := "asset/css/themes/default/assets/fonts/icons.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetHtmlButtonsTmpl reads file data from disk. It returns an error on failure.
func assetHtmlButtonsTmpl() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/html/buttons.tmpl"
	name := "asset/html/buttons.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetHtmlDevTmpl reads file data from disk. It returns an error on failure.
func assetHtmlDevTmpl() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/html/dev.tmpl"
	name := "asset/html/dev.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetHtmlIndexTmpl reads file data from disk. It returns an error on failure.
func assetHtmlIndexTmpl() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/html/index.tmpl"
	name := "asset/html/index.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetHtmlViewerTmpl reads file data from disk. It returns an error on failure.
func assetHtmlViewerTmpl() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/html/viewer.tmpl"
	name := "asset/html/viewer.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsGltfloaderJs reads file data from disk. It returns an error on failure.
func assetJsGltfloaderJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/GLTFLoader.js"
	name := "asset/js/GLTFLoader.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsOrbitcontrolsJs reads file data from disk. It returns an error on failure.
func assetJsOrbitcontrolsJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/OrbitControls.js"
	name := "asset/js/OrbitControls.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsAppJs reads file data from disk. It returns an error on failure.
func assetJsAppJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/app.js"
	name := "asset/js/app.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsAxiosMinJs reads file data from disk. It returns an error on failure.
func assetJsAxiosMinJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/axios.min.js"
	name := "asset/js/axios.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsComponentsJs reads file data from disk. It returns an error on failure.
func assetJsComponentsJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/components.js"
	name := "asset/js/components.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsCqpartsviewerJs reads file data from disk. It returns an error on failure.
func assetJsCqpartsviewerJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/cqpartsViewer.js"
	name := "asset/js/cqpartsViewer.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsSaverJs reads file data from disk. It returns an error on failure.
func assetJsSaverJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/saver.js"
	name := "asset/js/saver.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsSemanticUiVueMinJs reads file data from disk. It returns an error on failure.
func assetJsSemanticUiVueMinJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/semantic-ui-vue.min.js"
	name := "asset/js/semantic-ui-vue.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsThreeMinJs reads file data from disk. It returns an error on failure.
func assetJsThreeMinJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/three.min.js"
	name := "asset/js/three.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsVueJs reads file data from disk. It returns an error on failure.
func assetJsVueJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/vue.js"
	name := "asset/js/vue.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetJsVuexJs reads file data from disk. It returns an error on failure.
func assetJsVuexJs() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/js/vuex.js"
	name := "asset/js/vuex.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetVueAppVue reads file data from disk. It returns an error on failure.
func assetVueAppVue() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/vue/App.vue"
	name := "asset/vue/App.vue"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetVueCardVue reads file data from disk. It returns an error on failure.
func assetVueCardVue() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/vue/card.vue"
	name := "asset/vue/card.vue"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetVueModelVue reads file data from disk. It returns an error on failure.
func assetVueModelVue() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/vue/model.vue"
	name := "asset/vue/model.vue"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetVueNotifVue reads file data from disk. It returns an error on failure.
func assetVueNotifVue() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/vue/notif.vue"
	name := "asset/vue/notif.vue"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetVueStepsVue reads file data from disk. It returns an error on failure.
func assetVueStepsVue() (*asset, error) {
	path := "/opt/cqparts-server/src/server/asset/vue/steps.vue"
	name := "asset/vue/steps.vue"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"asset/css/extra.css": assetCssExtraCss,
	"asset/css/semantic.min.css": assetCssSemanticMinCss,
	"asset/css/themes/default/assets/fonts/icons.eot": assetCssThemesDefaultAssetsFontsIconsEot,
	"asset/css/themes/default/assets/fonts/icons.svg": assetCssThemesDefaultAssetsFontsIconsSvg,
	"asset/css/themes/default/assets/fonts/icons.ttf": assetCssThemesDefaultAssetsFontsIconsTtf,
	"asset/css/themes/default/assets/fonts/icons.woff": assetCssThemesDefaultAssetsFontsIconsWoff,
	"asset/css/themes/default/assets/fonts/icons.woff2": assetCssThemesDefaultAssetsFontsIconsWoff2,
	"asset/html/buttons.tmpl": assetHtmlButtonsTmpl,
	"asset/html/dev.tmpl": assetHtmlDevTmpl,
	"asset/html/index.tmpl": assetHtmlIndexTmpl,
	"asset/html/viewer.tmpl": assetHtmlViewerTmpl,
	"asset/js/GLTFLoader.js": assetJsGltfloaderJs,
	"asset/js/OrbitControls.js": assetJsOrbitcontrolsJs,
	"asset/js/app.js": assetJsAppJs,
	"asset/js/axios.min.js": assetJsAxiosMinJs,
	"asset/js/components.js": assetJsComponentsJs,
	"asset/js/cqpartsViewer.js": assetJsCqpartsviewerJs,
	"asset/js/saver.js": assetJsSaverJs,
	"asset/js/semantic-ui-vue.min.js": assetJsSemanticUiVueMinJs,
	"asset/js/three.min.js": assetJsThreeMinJs,
	"asset/js/vue.js": assetJsVueJs,
	"asset/js/vuex.js": assetJsVuexJs,
	"asset/vue/App.vue": assetVueAppVue,
	"asset/vue/card.vue": assetVueCardVue,
	"asset/vue/model.vue": assetVueModelVue,
	"asset/vue/notif.vue": assetVueNotifVue,
	"asset/vue/steps.vue": assetVueStepsVue,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"asset": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"extra.css": &bintree{assetCssExtraCss, map[string]*bintree{}},
			"semantic.min.css": &bintree{assetCssSemanticMinCss, map[string]*bintree{}},
			"themes": &bintree{nil, map[string]*bintree{
				"default": &bintree{nil, map[string]*bintree{
					"assets": &bintree{nil, map[string]*bintree{
						"fonts": &bintree{nil, map[string]*bintree{
							"icons.eot": &bintree{assetCssThemesDefaultAssetsFontsIconsEot, map[string]*bintree{}},
							"icons.svg": &bintree{assetCssThemesDefaultAssetsFontsIconsSvg, map[string]*bintree{}},
							"icons.ttf": &bintree{assetCssThemesDefaultAssetsFontsIconsTtf, map[string]*bintree{}},
							"icons.woff": &bintree{assetCssThemesDefaultAssetsFontsIconsWoff, map[string]*bintree{}},
							"icons.woff2": &bintree{assetCssThemesDefaultAssetsFontsIconsWoff2, map[string]*bintree{}},
						}},
					}},
				}},
			}},
		}},
		"html": &bintree{nil, map[string]*bintree{
			"buttons.tmpl": &bintree{assetHtmlButtonsTmpl, map[string]*bintree{}},
			"dev.tmpl": &bintree{assetHtmlDevTmpl, map[string]*bintree{}},
			"index.tmpl": &bintree{assetHtmlIndexTmpl, map[string]*bintree{}},
			"viewer.tmpl": &bintree{assetHtmlViewerTmpl, map[string]*bintree{}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"GLTFLoader.js": &bintree{assetJsGltfloaderJs, map[string]*bintree{}},
			"OrbitControls.js": &bintree{assetJsOrbitcontrolsJs, map[string]*bintree{}},
			"app.js": &bintree{assetJsAppJs, map[string]*bintree{}},
			"axios.min.js": &bintree{assetJsAxiosMinJs, map[string]*bintree{}},
			"components.js": &bintree{assetJsComponentsJs, map[string]*bintree{}},
			"cqpartsViewer.js": &bintree{assetJsCqpartsviewerJs, map[string]*bintree{}},
			"saver.js": &bintree{assetJsSaverJs, map[string]*bintree{}},
			"semantic-ui-vue.min.js": &bintree{assetJsSemanticUiVueMinJs, map[string]*bintree{}},
			"three.min.js": &bintree{assetJsThreeMinJs, map[string]*bintree{}},
			"vue.js": &bintree{assetJsVueJs, map[string]*bintree{}},
			"vuex.js": &bintree{assetJsVuexJs, map[string]*bintree{}},
		}},
		"vue": &bintree{nil, map[string]*bintree{
			"App.vue": &bintree{assetVueAppVue, map[string]*bintree{}},
			"card.vue": &bintree{assetVueCardVue, map[string]*bintree{}},
			"model.vue": &bintree{assetVueModelVue, map[string]*bintree{}},
			"notif.vue": &bintree{assetVueNotifVue, map[string]*bintree{}},
			"steps.vue": &bintree{assetVueStepsVue, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

