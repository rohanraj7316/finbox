# echo "**** **** START GENERATING DOCS **** ****"

OUTPUT=$(rm -rf ./../docs)
echo "${OUTPUT}"

OUTPUT=$(rm -rf ./../pkg)
echo "${OUTPUT}"

OUTPUT=$(buf format -w)
echo "${OUTPUT}"

OUTPUT=$(buf generate)
echo "${OUTPUT}"

# echo "**** **** END GENERATING DOCS **** ****"