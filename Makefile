.PHONY: create_question
create_question:
	@read -p "Enter question folder (e.g. question6): " folder; \
	./create_question.sh $$folder

