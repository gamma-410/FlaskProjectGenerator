#!bin/shell

mkdir newProject
cd newProject

python3 -m venv venv
source venv/bin/activate

pip install flask

mkdir templates
mkdir static
cp ../src/app.py  app.py
cp ../src/templates/index.html  templates/index.html
cp ../src/templates/base.html  templates/base.html
cp ../src/static/style.css  static/style.css
cp ../src/static/main.js  static/main.js

deactivate

echo   
echo 処理が完了しました。
echo プロジェクトのディレクトリに移動して source venv/bin/activate を実行してください。
echo python3 app.py でサーバーが起動します。
echo


