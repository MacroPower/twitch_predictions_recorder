package twitchtest_test

import (
	"bytes"
	"testing"

	"github.com/MacroPower/twitch_predictions_recorder/internal/event"
	"github.com/MacroPower/twitch_predictions_recorder/internal/twitchtest"

	"github.com/stretchr/testify/require"
)

func TestReplay(t *testing.T) {
	t.Parallel()

	b := bytes.Buffer{}

	b.WriteString(`{"type":"event-updated","data":{"timestamp":"2023-05-10T04:52:33.8848932Z","event":{"id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","channel_id":"121059319","created_at":"2023-05-10T03:19:40.218367915Z","created_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"ended_at":"2023-05-10T04:52:33.81954282Z","ended_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"locked_at":"2023-05-10T03:29:39.369374545Z","locked_by":{"type":"","user_id":"","user_display_name":"","extension_client_id":null},"outcomes":[{"id":"e4776805-aafc-427c-8022-531cbe7740d4","color":"BLUE","title":"Yes","total_points":19339491,"total_users":753,"top_predictors":[{"id":"ab7547b807db75b85f12544bb8869b08be7fff256cb108279ef206153e779a3f","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.888319787Z","updated_at":"2023-05-10T03:19:57.888319787Z","user_id":"125741983","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"J3RlN"},{"id":"9828d04f4d0239a3ad7e04874b614355f1185a092bb15d95040dc2a881cf785e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.77230864Z","updated_at":"2023-05-10T03:20:02.77230864Z","user_id":"53468991","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"IamLuck_"},{"id":"88c9bd3f4b1eb2d48bfd68918ac08137d9efa30de66dc0c69410eb143c6ace9e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:09.591704954Z","updated_at":"2023-05-10T03:20:09.591704954Z","user_id":"25378114","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"Terkeal"},{"id":"36ededc69923b67c717c2c6976d13ee7506467cf6fc04cef88125ffc192ffb9c","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:12.418755506Z","updated_at":"2023-05-10T03:20:12.418755506Z","user_id":"55238137","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"PokerBluffalo"},{"id":"f227a2ab6f23fe89470d832a813ea9e0dbe1de709138a55d4f63ddcdebf4b1da","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:23.728259735Z","updated_at":"2023-05-10T03:20:23.728259735Z","user_id":"174052226","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"shvkedown"},{"id":"2339cd6b4aae889aee4bc583b564fb73177577066142ae241a870cd0499607c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:33.260680127Z","updated_at":"2023-05-10T03:20:33.260680127Z","user_id":"154392346","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"b_riles"},{"id":"6128f2c7f4b31bdee7375e2d26692e5e9839a461dd70697990674ed4f37cef64","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:40.164055005Z","updated_at":"2023-05-10T03:20:40.164055005Z","user_id":"28430287","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"deciduous"},{"id":"e326d11400cab7cf6408d588932cade0d963b30f0dfc853f25db834db16863c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:17.546441543Z","updated_at":"2023-05-10T03:20:43.285755609Z","user_id":"57218750","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"mitchocity"},{"id":"0d09b82e7d3d6ea9d4136c1d885913d3f8db18a9527232661f3b40be108cbd08","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:46.993717981Z","updated_at":"2023-05-10T03:20:46.993717981Z","user_id":"37152061","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"Klubbah"},{"id":"003c803eb81e18f92634c3246b9dea80bd07511bcb50880db42449f8aef62912","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:55.433092662Z","updated_at":"2023-05-10T03:20:55.433092662Z","user_id":"416310969","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"nikomiri"}],"badge":{"version":"blue-1","set_id":"predictions"}},{"id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","color":"PINK","title":"No","total_points":24237598,"total_users":675,"top_predictors":[{"id":"3723babed8ac3a7309c8ff1a6b117382b680449a1ba08530f66571991f857ccd","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:55.273406806Z","updated_at":"2023-05-10T03:19:55.273406806Z","user_id":"84880542","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"shwooders"},{"id":"d2f263208ef3cdbf2f42437ece57542f062f32995a8afd3010169bf28fdbf363","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:56.600210772Z","updated_at":"2023-05-10T03:19:56.600210772Z","user_id":"732229297","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"popesiclestick"},{"id":"3271ba72ad5ae7b09178042c371832c76b9920ad4aa16a9637b7fb964feb088b","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.131683952Z","updated_at":"2023-05-10T03:19:57.131683952Z","user_id":"106332976","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"clapmygorillacheeks"},{"id":"32a506ecc68ba7514f807c6ae0deacf40fe88dd868657ece9f7ec9f84dac4a31","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.581126219Z","updated_at":"2023-05-10T03:19:57.581126219Z","user_id":"112868955","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"prismbreak_"},{"id":"41e56b28ec344506fec00ab32bc0f6e2f59a047afa9234ef520467ac0b3fc5cf","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:00.643877743Z","updated_at":"2023-05-10T03:20:00.643877743Z","user_id":"32185979","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"SilentFrost"},{"id":"2d97e2a7aa617aea6d52b541d4b6b951e56a9b9bdc60feb3e22a515a9d28fa65","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.723661021Z","updated_at":"2023-05-10T03:20:02.723661021Z","user_id":"30295435","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"kia176"},{"id":"6bc90dfd8c70aecb45184008414261e3feed0a3bf2f7ba7d8dc83ddd3e658350","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.759915231Z","updated_at":"2023-05-10T03:20:06.759915231Z","user_id":"154052650","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"tcssj2gohan"},{"id":"5cd4be59e1c2b622c73feb76e48a926ae4c523e69434e2455d8e107e4362144d","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.798455178Z","updated_at":"2023-05-10T03:20:06.798455178Z","user_id":"67112113","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"zzesterino"},{"id":"d859f0dd31cbc2759e461738e26dd379c8ee0eeabf15f78663d2f8b432c0e6c9","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:08.665964717Z","updated_at":"2023-05-10T03:20:08.665964717Z","user_id":"46677495","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"vinnsonn"},{"id":"a5c7f2ba1962d886455152711372c416a01499c5671976fd92193afde6687670","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:18.955531551Z","updated_at":"2023-05-10T03:20:18.955531551Z","user_id":"46540524","result":{"type":"","points_won":0,"is_acknowledged":false},"user_display_name":"chrisb960"}],"badge":{"version":"pink-2","set_id":"predictions"}}],"prediction_window_seconds":600,"status":"RESOLVE_PENDING","title":"Will moon Survive the Third Region?","winning_outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939"}}}` + "\n")
	b.WriteString(`{"type":"event-updated","data":{"timestamp":"2023-05-10T04:52:34.064523572Z","event":{"id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","channel_id":"121059319","created_at":"2023-05-10T03:19:40.218367915Z","created_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"ended_at":"2023-05-10T04:52:33.81954282Z","ended_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"locked_at":"2023-05-10T03:29:39.369374545Z","locked_by":{"type":"","user_id":"","user_display_name":"","extension_client_id":null},"outcomes":[{"id":"e4776805-aafc-427c-8022-531cbe7740d4","color":"BLUE","title":"Yes","total_points":19439521,"total_users":754,"top_predictors":[{"id":"ab7547b807db75b85f12544bb8869b08be7fff256cb108279ef206153e779a3f","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.888319787Z","updated_at":"2023-05-10T03:19:57.888319787Z","user_id":"125741983","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"J3RlN"},{"id":"9828d04f4d0239a3ad7e04874b614355f1185a092bb15d95040dc2a881cf785e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.77230864Z","updated_at":"2023-05-10T03:20:02.77230864Z","user_id":"53468991","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"IamLuck_"},{"id":"88c9bd3f4b1eb2d48bfd68918ac08137d9efa30de66dc0c69410eb143c6ace9e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:09.591704954Z","updated_at":"2023-05-10T03:20:09.591704954Z","user_id":"25378114","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"Terkeal"},{"id":"36ededc69923b67c717c2c6976d13ee7506467cf6fc04cef88125ffc192ffb9c","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:12.418755506Z","updated_at":"2023-05-10T03:20:12.418755506Z","user_id":"55238137","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"PokerBluffalo"},{"id":"f227a2ab6f23fe89470d832a813ea9e0dbe1de709138a55d4f63ddcdebf4b1da","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:23.728259735Z","updated_at":"2023-05-10T03:20:23.728259735Z","user_id":"174052226","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"shvkedown"},{"id":"2339cd6b4aae889aee4bc583b564fb73177577066142ae241a870cd0499607c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:33.260680127Z","updated_at":"2023-05-10T03:20:33.260680127Z","user_id":"154392346","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"b_riles"},{"id":"6128f2c7f4b31bdee7375e2d26692e5e9839a461dd70697990674ed4f37cef64","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:40.164055005Z","updated_at":"2023-05-10T03:20:40.164055005Z","user_id":"28430287","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"deciduous"},{"id":"e326d11400cab7cf6408d588932cade0d963b30f0dfc853f25db834db16863c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:17.546441543Z","updated_at":"2023-05-10T03:20:43.285755609Z","user_id":"57218750","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"mitchocity"},{"id":"0d09b82e7d3d6ea9d4136c1d885913d3f8db18a9527232661f3b40be108cbd08","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:46.993717981Z","updated_at":"2023-05-10T03:20:46.993717981Z","user_id":"37152061","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"Klubbah"},{"id":"003c803eb81e18f92634c3246b9dea80bd07511bcb50880db42449f8aef62912","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:55.433092662Z","updated_at":"2023-05-10T03:20:55.433092662Z","user_id":"416310969","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"nikomiri"}],"badge":{"version":"blue-1","set_id":"predictions"}},{"id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","color":"PINK","title":"No","total_points":24238598,"total_users":676,"top_predictors":[{"id":"3723babed8ac3a7309c8ff1a6b117382b680449a1ba08530f66571991f857ccd","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:55.273406806Z","updated_at":"2023-05-10T03:19:55.273406806Z","user_id":"84880542","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"shwooders"},{"id":"d2f263208ef3cdbf2f42437ece57542f062f32995a8afd3010169bf28fdbf363","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:56.600210772Z","updated_at":"2023-05-10T03:19:56.600210772Z","user_id":"732229297","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"popesiclestick"},{"id":"3271ba72ad5ae7b09178042c371832c76b9920ad4aa16a9637b7fb964feb088b","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.131683952Z","updated_at":"2023-05-10T03:19:57.131683952Z","user_id":"106332976","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"clapmygorillacheeks"},{"id":"32a506ecc68ba7514f807c6ae0deacf40fe88dd868657ece9f7ec9f84dac4a31","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.581126219Z","updated_at":"2023-05-10T03:19:57.581126219Z","user_id":"112868955","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"prismbreak_"},{"id":"41e56b28ec344506fec00ab32bc0f6e2f59a047afa9234ef520467ac0b3fc5cf","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:00.643877743Z","updated_at":"2023-05-10T03:20:00.643877743Z","user_id":"32185979","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"SilentFrost"},{"id":"2d97e2a7aa617aea6d52b541d4b6b951e56a9b9bdc60feb3e22a515a9d28fa65","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.723661021Z","updated_at":"2023-05-10T03:20:02.723661021Z","user_id":"30295435","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"kia176"},{"id":"6bc90dfd8c70aecb45184008414261e3feed0a3bf2f7ba7d8dc83ddd3e658350","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.759915231Z","updated_at":"2023-05-10T03:20:06.759915231Z","user_id":"154052650","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"tcssj2gohan"},{"id":"5cd4be59e1c2b622c73feb76e48a926ae4c523e69434e2455d8e107e4362144d","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.798455178Z","updated_at":"2023-05-10T03:20:06.798455178Z","user_id":"67112113","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"zzesterino"},{"id":"d859f0dd31cbc2759e461738e26dd379c8ee0eeabf15f78663d2f8b432c0e6c9","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:08.665964717Z","updated_at":"2023-05-10T03:20:08.665964717Z","user_id":"46677495","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"vinnsonn"},{"id":"a5c7f2ba1962d886455152711372c416a01499c5671976fd92193afde6687670","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:18.955531551Z","updated_at":"2023-05-10T03:20:18.955531551Z","user_id":"46540524","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"chrisb960"}],"badge":{"version":"pink-2","set_id":"predictions"}}],"prediction_window_seconds":600,"status":"RESOLVE_PENDING","title":"Will moon Survive the Third Region?","winning_outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939"}}}` + "\n")
	b.WriteString(`{"type":"event-updated","data":{"timestamp":"2023-05-10T04:52:39.941885142Z","event":{"id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","channel_id":"121059319","created_at":"2023-05-10T03:19:40.218367915Z","created_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"ended_at":"2023-05-10T04:52:33.81954282Z","ended_by":{"type":"USER","user_id":"95987079","user_display_name":"EnDecc","extension_client_id":null},"locked_at":"2023-05-10T03:29:39.369374545Z","locked_by":{"type":"","user_id":"","user_display_name":"","extension_client_id":null},"outcomes":[{"id":"e4776805-aafc-427c-8022-531cbe7740d4","color":"BLUE","title":"Yes","total_points":19439521,"total_users":754,"top_predictors":[{"id":"ab7547b807db75b85f12544bb8869b08be7fff256cb108279ef206153e779a3f","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.888319787Z","updated_at":"2023-05-10T03:19:57.888319787Z","user_id":"125741983","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"J3RlN"},{"id":"9828d04f4d0239a3ad7e04874b614355f1185a092bb15d95040dc2a881cf785e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.77230864Z","updated_at":"2023-05-10T03:20:02.77230864Z","user_id":"53468991","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"IamLuck_"},{"id":"88c9bd3f4b1eb2d48bfd68918ac08137d9efa30de66dc0c69410eb143c6ace9e","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:09.591704954Z","updated_at":"2023-05-10T03:20:09.591704954Z","user_id":"25378114","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"Terkeal"},{"id":"36ededc69923b67c717c2c6976d13ee7506467cf6fc04cef88125ffc192ffb9c","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:12.418755506Z","updated_at":"2023-05-10T03:20:12.418755506Z","user_id":"55238137","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"PokerBluffalo"},{"id":"f227a2ab6f23fe89470d832a813ea9e0dbe1de709138a55d4f63ddcdebf4b1da","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:23.728259735Z","updated_at":"2023-05-10T03:20:23.728259735Z","user_id":"174052226","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"shvkedown"},{"id":"2339cd6b4aae889aee4bc583b564fb73177577066142ae241a870cd0499607c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:33.260680127Z","updated_at":"2023-05-10T03:20:33.260680127Z","user_id":"154392346","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"b_riles"},{"id":"6128f2c7f4b31bdee7375e2d26692e5e9839a461dd70697990674ed4f37cef64","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:40.164055005Z","updated_at":"2023-05-10T03:20:40.164055005Z","user_id":"28430287","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"deciduous"},{"id":"e326d11400cab7cf6408d588932cade0d963b30f0dfc853f25db834db16863c6","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:17.546441543Z","updated_at":"2023-05-10T03:20:43.285755609Z","user_id":"57218750","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"mitchocity"},{"id":"0d09b82e7d3d6ea9d4136c1d885913d3f8db18a9527232661f3b40be108cbd08","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:46.993717981Z","updated_at":"2023-05-10T03:20:46.993717981Z","user_id":"37152061","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"Klubbah"},{"id":"003c803eb81e18f92634c3246b9dea80bd07511bcb50880db42449f8aef62912","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"e4776805-aafc-427c-8022-531cbe7740d4","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:55.433092662Z","updated_at":"2023-05-10T03:20:55.433092662Z","user_id":"416310969","result":{"type":"LOSE","points_won":0,"is_acknowledged":false},"user_display_name":"nikomiri"}],"badge":{"version":"blue-1","set_id":"predictions"}},{"id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","color":"PINK","title":"No","total_points":24238598,"total_users":676,"top_predictors":[{"id":"3723babed8ac3a7309c8ff1a6b117382b680449a1ba08530f66571991f857ccd","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:55.273406806Z","updated_at":"2023-05-10T03:19:55.273406806Z","user_id":"84880542","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"shwooders"},{"id":"d2f263208ef3cdbf2f42437ece57542f062f32995a8afd3010169bf28fdbf363","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:56.600210772Z","updated_at":"2023-05-10T03:19:56.600210772Z","user_id":"732229297","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"popesiclestick"},{"id":"3271ba72ad5ae7b09178042c371832c76b9920ad4aa16a9637b7fb964feb088b","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.131683952Z","updated_at":"2023-05-10T03:19:57.131683952Z","user_id":"106332976","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"clapmygorillacheeks"},{"id":"32a506ecc68ba7514f807c6ae0deacf40fe88dd868657ece9f7ec9f84dac4a31","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:19:57.581126219Z","updated_at":"2023-05-10T03:19:57.581126219Z","user_id":"112868955","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"prismbreak_"},{"id":"41e56b28ec344506fec00ab32bc0f6e2f59a047afa9234ef520467ac0b3fc5cf","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:00.643877743Z","updated_at":"2023-05-10T03:20:00.643877743Z","user_id":"32185979","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"SilentFrost"},{"id":"2d97e2a7aa617aea6d52b541d4b6b951e56a9b9bdc60feb3e22a515a9d28fa65","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:02.723661021Z","updated_at":"2023-05-10T03:20:02.723661021Z","user_id":"30295435","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"kia176"},{"id":"6bc90dfd8c70aecb45184008414261e3feed0a3bf2f7ba7d8dc83ddd3e658350","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.759915231Z","updated_at":"2023-05-10T03:20:06.759915231Z","user_id":"154052650","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"tcssj2gohan"},{"id":"5cd4be59e1c2b622c73feb76e48a926ae4c523e69434e2455d8e107e4362144d","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:06.798455178Z","updated_at":"2023-05-10T03:20:06.798455178Z","user_id":"67112113","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"zzesterino"},{"id":"d859f0dd31cbc2759e461738e26dd379c8ee0eeabf15f78663d2f8b432c0e6c9","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:08.665964717Z","updated_at":"2023-05-10T03:20:08.665964717Z","user_id":"46677495","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"vinnsonn"},{"id":"a5c7f2ba1962d886455152711372c416a01499c5671976fd92193afde6687670","event_id":"f5ff8c36-a3b3-4de6-97b1-657063e0f259","outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939","channel_id":"121059319","points":250000,"predicted_at":"2023-05-10T03:20:18.955531551Z","updated_at":"2023-05-10T03:20:18.955531551Z","user_id":"46540524","result":{"type":"WIN","points_won":450502,"is_acknowledged":false},"user_display_name":"chrisb960"}],"badge":{"version":"pink-2","set_id":"predictions"}}],"prediction_window_seconds":600,"status":"RESOLVED","title":"Will moon Survive the Third Region?","winning_outcome_id":"ea190e0b-6110-4f03-8217-d75d0ddfb939"}}}` + "\n")

	i := 0
	lt := twitchtest.NewTestListener(&b)
	err := lt.Listen(func(e event.Event) error {
		require.Equal(t, "event-updated", e.EventStates[0].Type)
		switch i {
		case 0:
		case 1:
			require.Equal(t, "RESOLVE_PENDING", e.EventStates[0].Status)
		case 2:
			require.Equal(t, "RESOLVED", e.EventStates[0].Status)
		}
		i++

		return nil
	})
	require.NoError(t, err)
}
